package db

import (
	"bibi/pkg/constants"
	"context"
	"fmt"
	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
	"strconv"
)

func CreateCollection(ctx context.Context) error {
	exist, err := MilvusCli.HasCollection(ctx, constants.AvatarMilvusName)
	if err != nil {
		return err
	}

	if exist {
		// colletion exist,不作为错误
		return nil
	}

	schema := &entity.Schema{
		CollectionName: constants.AvatarMilvusName,
		Description:    "search for avatar",
		Fields: []*entity.Field{
			{
				Name:       "id",
				DataType:   entity.FieldTypeInt64,
				PrimaryKey: true,
				AutoID:     true,
			},
			{
				Name:       "uid",
				DataType:   entity.FieldTypeInt64,
				PrimaryKey: false,
				AutoID:     false,
			},
			{
				Name:     "avatar",
				DataType: entity.FieldTypeFloatVector,
				TypeParams: map[string]string{
					"dim": strconv.Itoa(constants.Dim),
				},
			},
		},
		EnableDynamicField: true,
	}

	if err := MilvusCli.CreateCollection(ctx, schema, 3); err != nil {
		return err
	}

	idx, err := entity.NewIndexIvfFlat( // NewIndex func
		entity.L2,     // metricType
		constants.Dim, // ConstructParams
	)
	if err != nil {
		return err
	}

	if err = MilvusCli.CreateIndex(
		ctx,                        // ctx
		constants.AvatarMilvusName, // CollectionName
		"avatar",                   // fieldName
		idx,                        // entity.Index
		false,                      // async
	); err != nil {
		return err
	}
	return nil
}

func InsertData(ctx context.Context, vector []float32, uid int64) error {
	uidC := make([]int64, 0)
	avatarC := make([][]float32, 0)
	uidC = append(uidC, uid)
	avatarC = append(avatarC, vector)

	uidColumn := entity.NewColumnInt64("uid", uidC)
	//fmt.Println(uidC)
	avatarColumn := entity.NewColumnFloatVector("avatar", constants.Dim, avatarC)

	if _, err := MilvusCli.Insert(
		ctx,
		constants.AvatarMilvusName,
		"",
		uidColumn,
		avatarColumn,
	); err != nil {
		return err
	}

	return nil
}

func Search(ctx context.Context, vector []float32) ([]int64, error) {

	err := MilvusCli.LoadCollection(ctx, constants.AvatarMilvusName, false)
	if err != nil {
		return nil, err
	}

	sp, _ := entity.NewIndexIvfFlatSearchParam(
		50,
	)

	opt := client.SearchQueryOptionFunc(func(option *client.SearchQueryOption) {
		option.Limit = 3
		option.Offset = 0
		option.ConsistencyLevel = entity.ClStrong //
		option.IgnoreGrowing = false
	})
	searchResult, err := MilvusCli.Search(
		ctx,                        // ctx
		constants.AvatarMilvusName, // CollectionName
		[]string{},                 // partitionNames
		"",                         // expr
		[]string{"uid"},            // outputFields
		[]entity.Vector{entity.FloatVector(vector)}, // vectors
		"avatar",  // vectorField
		entity.L2, // metricType
		10,        // topK
		sp,        // searchParams
		opt,
	)

	uidList := make([]int64, 0)
	if err != nil {
		return nil, err
	}
	for _, v := range searchResult {
		for i := 0; i < v.ResultCount; i++ {
			fmt.Println(searchResult)
			id, err := v.Fields.GetColumn("uid").GetAsInt64(i)
			if err != nil {
				continue
			}
			uidList = append(uidList, id)
		}
	}

	//err = MilvusCli.ReleaseCollection(ctx, constants.AvatarMilvusName)
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Println("uid", uidList)
	return uidList, nil
}
