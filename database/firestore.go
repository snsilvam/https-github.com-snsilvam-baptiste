package database

import (
	"context"
	"fmt"

	"baptiste.com/models"
	"cloud.google.com/go/firestore"
)

type FirestoreRepository struct {
	client *firestore.Client
}

func NewClientFirestore(ctx context.Context, projectID string) (*FirestoreRepository, error) {
	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		return nil, err
	}

	return &FirestoreRepository{client}, nil
}

func (f *FirestoreRepository) InsertMonthlyExpenses(ctx context.Context, monthlyExpenses *models.MonthlyExpensesModelInsert) error {
	collectionMonthlyExpenses := f.client.Collection("monthlyExpensesModel")

	wr, err := collectionMonthlyExpenses.NewDoc().Create(ctx, monthlyExpenses)
	if err != nil {
		fmt.Println("error al intentar crear el documento en firestore:", err)
		return err
	}

	fmt.Println("El documento se creo con exito ☻", wr)

	return nil
}

func (f *FirestoreRepository) GetMonthlyExpense(ctx context.Context, id string) (*models.MonthlyExpensesModel, error) {
	var monthlyExpensesModel models.MonthlyExpensesModel

	doc := f.client.Doc("monthlyExpensesModel/" + id)

	docsnap, err := doc.Get(ctx)
	if err != nil {
		return nil, err
	}

	if err = docsnap.DataTo(&monthlyExpensesModel); err != nil {
		return nil, err
	}

	monthlyExpensesModel.ID = docsnap.Ref.ID

	return &monthlyExpensesModel, nil
}

func (f *FirestoreRepository) UpdateMonthlyExpense(ctx context.Context, monthlyExpense *models.MonthlyExpensesModelUpdate) error {
	doc := f.client.Doc("monthlyExpensesModel/" + monthlyExpense.ID)

	_, err := doc.Update(ctx, []firestore.Update{{Path: "NameFixedExpense", Value: monthlyExpense.NameFixedExpense}, {Path: "DueDate", Value: monthlyExpense.DueDate}, {Path: "Status", Value: monthlyExpense.Status}})
	if err != nil {
		return err
	}
	return nil
}
