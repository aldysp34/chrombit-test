package usecase

import (
	"aldysp34/chrombit-test/apperror"
	"aldysp34/chrombit-test/dto"
	"context"
)

type BlogUsecase interface {
	GetBlogs(context.Context) []dto.BlogResponse
	GetBlogByID(context.Context, int) (*dto.BlogResponse, error)
	EditBlog(context.Context, dto.BlogRequest) (*dto.BlogResponse, error)
	DeleteBlog(context.Context, int) error
	CreateBlog(context.Context, dto.BlogRequest) (*dto.BlogResponse, error)
}

type BlogDummyData struct {
	ID    int
	Title string
	Body  string
	Slug  string
}

var blogData []BlogDummyData

type blogUsecase struct{}

func NewBlogUsecase() BlogUsecase {
	blogData = append(blogData, BlogDummyData{1, "Introduction to Go Programming", "Go is a powerful programming language designed for simplicity and efficiency.", "go-programming-intro"})
	blogData = append(blogData, BlogDummyData{2, "Building RESTful APIs with Node.js", "Learn how to create RESTful APIs using Node.js and Express.", "nodejs-restful-apis"})
	blogData = append(blogData, BlogDummyData{3, "Python for Data Science: A Comprehensive Guide", "Explore the world of data science with Python and its rich ecosystem of libraries.", "python-data-science-guide"})
	blogData = append(blogData, BlogDummyData{4, "Web Development with React.js", "Build modern and interactive web applications with the React.js library.", "reactjs-web-development"})
	blogData = append(blogData, BlogDummyData{5, "Getting Started with Docker Containers", "An introduction to Docker and containerization for scalable and portable applications.", "docker-getting-started"})
	blogData = append(blogData, BlogDummyData{6, "Machine Learning Fundamentals", "Understand the basics of machine learning and its applications in various domains.", "machine-learning-fundamentals"})
	blogData = append(blogData, BlogDummyData{7, "Java Programming Best Practices", "Follow best practices to write clean, efficient, and maintainable Java code.", "java-best-practices"})
	blogData = append(blogData, BlogDummyData{8, "Introduction to Cybersecurity", "Learn the fundamentals of cybersecurity to protect your digital assets and data.", "cybersecurity-intro"})
	blogData = append(blogData, BlogDummyData{9, "iOS App Development with Swift", "Create native iOS applications using the Swift programming language and Xcode.", "swift-ios-app-development"})
	blogData = append(blogData, BlogDummyData{10, "Database Design and Optimization Techniques", "Design and optimize databases for improved performance and reliability.", "database-design-optimization"})
	return &blogUsecase{}
}

func (bu *blogUsecase) GetBlogs(ctx context.Context) []dto.BlogResponse {
	var response []dto.BlogResponse
	for _, v := range blogData {
		resp := dto.BlogResponse{
			ID:    v.ID,
			Title: v.Title,
			Body:  v.Body,
			Slug:  v.Slug,
		}

		response = append(response, resp)
	}

	return response
}

func (bu *blogUsecase) GetBlogByID(ctx context.Context, id int) (*dto.BlogResponse, error) {
	for _, v := range blogData {
		if v.ID == id {
			return &dto.BlogResponse{
				ID:    v.ID,
				Title: v.Title,
				Body:  v.Body,
				Slug:  v.Slug,
			}, nil
		}
	}

	return nil, apperror.ErrBlogNotFound
}

func (bu *blogUsecase) EditBlog(ctx context.Context, req dto.BlogRequest) (*dto.BlogResponse, error) {
	for i, v := range blogData {
		if v.ID == req.ID {
			blogData[i].Body = req.Body
			blogData[i].Title = req.Title
			blogData[i].Slug = req.Slug

			return &dto.BlogResponse{
				ID:    blogData[i].ID,
				Title: blogData[i].Title,
				Body:  blogData[i].Body,
				Slug:  blogData[i].Slug,
			}, nil
		}
	}

	return nil, apperror.ErrBlogNotFound
}

func (bu *blogUsecase) CreateBlog(ctx context.Context, req dto.BlogRequest) (*dto.BlogResponse, error) {
	newData := BlogDummyData{
		ID:    req.ID,
		Body:  req.Body,
		Title: req.Title,
		Slug:  req.Slug,
	}

	for _, v := range blogData {
		if v.ID == newData.ID {
			return nil, apperror.ErrDuplicateBlog
		}
	}
	blogData = append(blogData, newData)

	return &dto.BlogResponse{
		ID:    newData.ID,
		Title: newData.Title,
		Body:  newData.Body,
		Slug:  newData.Slug,
	}, nil
}

func (bu *blogUsecase) DeleteBlog(ctx context.Context, id int) error {
	for i, v := range blogData {
		if v.ID == id {
			bu.deleteElement(i)

			return nil
		}
	}

	return apperror.ErrBlogNotFound

}

func (bu *blogUsecase) deleteElement(index int) {
	blogData = append(blogData[:index], blogData[index+1:]...)
}
