package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"sirclo/gql/entities"
	"sirclo/gql/entities/graph/model"
	"sirclo/gql/pkg/jwt"
	auth "sirclo/gql/repository/authmiddleware"
	"sirclo/gql/util/graph/generated"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	UserData := entities.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	ResponseData, err := r.userRepo.CreateUser(UserData)
	if err != nil {
		return nil, errors.New("failed create user")
	}

	UserResponseData := model.User{
		ID:    &ResponseData.Id,
		Name:  ResponseData.Name,
		Email: ResponseData.Email,
	}
	return &UserResponseData, nil
}

func (r *mutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.Book, error) {
	username := auth.ForContext(ctx)
	fmt.Println("ini username & ctx", username, ctx)
	if username == nil {
		return &model.Book{}, fmt.Errorf("unauthorized")
	}
	bookData := entities.Book{
		Name:      input.Name,
		Author:    input.Author,
		Publisher: input.Publisher,
		Year:      input.Year,
	}

	responseData, err := r.bookRepo.CreateBook(bookData)
	if err != nil {
		return nil, errors.New("failed create book")
	}

	bookResponseData := model.Book{
		ID:        &responseData.Id,
		Name:      responseData.Name,
		Author:    responseData.Author,
		Publisher: responseData.Publisher,
		Year:      responseData.Year,
	}

	return &bookResponseData, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input model.NewUser) (*model.User, error) {
	username := auth.ForContext(ctx)
	fmt.Println("ini username & ctx", username, ctx)
	if username == nil {
		return &model.User{}, fmt.Errorf("unauthorized")
	}
	user, err := r.userRepo.GetUser(id)
	if err != nil {
		return nil, errors.New("not found")
	}
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password

	updateUser, err := r.userRepo.UpdateUser(user)
	modelUU := model.User{
		Name:     updateUser.Name,
		Email:    updateUser.Email,
		Password: updateUser.Password,
	}
	return &modelUU, err
}

func (r *mutationResolver) UpdateBook(ctx context.Context, id int, input model.NewBook) (*model.Book, error) {
	username := auth.ForContext(ctx)
	fmt.Println("ini username & ctx", username, ctx)
	if username == nil {
		return &model.Book{}, fmt.Errorf("unauthorized")
	}
	book, err := r.bookRepo.GetBook(id)
	if err != nil {
		return nil, errors.New("not found")
	}
	book.Name = input.Name
	book.Author = input.Author
	book.Publisher = input.Publisher
	book.Year = input.Year

	UpdateBook, err := r.bookRepo.UpdateBook(book)
	modelUU := model.Book{
		Name:      UpdateBook.Name,
		Author:    UpdateBook.Author,
		Publisher: UpdateBook.Publisher,
		Year:      UpdateBook.Year,
	}
	return &modelUU, err
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*model.Message, error) {
	username := auth.ForContext(ctx)
	fmt.Println("ini username & ctx", username, ctx)
	if username == nil {
		return &model.Message{}, fmt.Errorf("unauthorized")
	}
	user, err := r.userRepo.GetUser(id)
	if err != nil {
		return nil, errors.New("not found")
	}
	_, err = r.userRepo.DeleteUser(user)
	if err != nil {
		return nil, err
	}

	return &model.Message{Code: 200, Message: "Delete User Success"}, nil
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id int) (*model.Message, error) {
	username := auth.ForContext(ctx)
	fmt.Println("ini username & ctx", username, ctx)
	if username == nil {
		return &model.Message{}, fmt.Errorf("unauthorized")
	}
	book, err := r.bookRepo.GetBook(id)
	if err != nil {
		return nil, errors.New("not found")
	}
	_, err = r.bookRepo.DeleteBook(book)
	if err != nil {
		return nil, err
	}

	return &model.Message{Code: 200, Message: "Delete Book Success"}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	username := auth.ForContext(ctx)
	fmt.Println("ini username & ctx", username, ctx)
	if username == nil {
		return []*model.User{}, fmt.Errorf("unauthorized")
	}
	fmt.Println(username)
	responseData, err := r.userRepo.GetUsers()
	fmt.Println(responseData)

	if err != nil {
		return nil, errors.New("not found")
	}

	UserResponseData := []*model.User{}

	for _, v := range responseData {
		convertID := v.Id
		UserResponseData = append(UserResponseData, &model.User{ID: &convertID, Name: v.Name, Email: v.Email})
	}

	return UserResponseData, nil
}

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	username := auth.ForContext(ctx)
	fmt.Println("ini username & ctx", username, ctx)
	if username == nil {
		return []*model.Book{}, fmt.Errorf("unauthorized")
	}
	responseData, err := r.bookRepo.GetBooks()
	if err != nil {
		return nil, errors.New("not found")
	}

	BookResponseData := []*model.Book{}

	for _, v := range responseData {
		BookResponseData = append(BookResponseData, &model.Book{ID: &v.Id, Name: v.Name, Author: v.Author, Publisher: v.Publisher, Year: v.Year})
	}
	return BookResponseData, nil
}

func (r *queryResolver) User(ctx context.Context, id *int) (*model.User, error) {
	responseData, err := r.userRepo.GetUser(*id)
	username := auth.ForContext(ctx)
	fmt.Println("ini username & ctx", username, ctx)
	if username == nil {
		return &model.User{}, fmt.Errorf("unauthorized")
	}
	// fmt.Errorf(responseData.Name)
	if err != nil {
		return nil, errors.New("not found")
	}
	modelData := model.User{
		ID:    &responseData.Id,
		Name:  responseData.Name,
		Email: responseData.Email,
	}
	return &modelData, nil
}

func (r *queryResolver) Login(ctx context.Context, name string, password string) (*model.LoginResponse, error) {
	correct := r.userRepo.Authenticate(name, password)
	if !correct {
		// 1
		return nil, fmt.Errorf("wrong username or password")
	}
	token, err := jwt.GenerateToken(name)
	if err != nil {
		return nil, err
	}
	modelLR := model.LoginResponse{
		Message: "Login Success",
		ID:      nil,
		Name:    nil,
		Token:   &token,
	}
	return &modelLR, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
