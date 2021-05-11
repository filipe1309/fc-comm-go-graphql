# DevOnTheRun Notes

> notes taken during the course

<!-- https://gitignore.io -->

```sh
go mod init go-graphql
go get github.com/99designs/gqlgen
go run github.com/99designs/gqlgen init

go get github.com/vektah/gqlparser/v2@v2.1.0
go get github.com/vektah/gqlparser/v2/ast@v2.1.0

go run server.go
http://localhost:8080/

```

http://localhost:8080/query

```graphql
query findCategories {
    categories {
        id
        name
        description
        courses {
            name
        }
    }
}

query findCourses {
    courses {
        id
        name
        description
        chapters {
            name
        }
        category {
            name
        }
    }
}

mutation createCategory {
    createCategory(input: { name: "PHP", description: "PHP is awesome" }) {
        id
        name
        description
    }
}

mutation createCourse {
    createCourse(
        input: {
            name: "Evo PHP"
            description: "PHP is awesome"
            categoryId: "T5577006791947779410"
        }
    ) {
        id
        name
        description
        category {
            id
            name
        }
    }
}

mutation createChapter {
    createChapter(
        input: { name: "Evo PHP - C1", courseId: "T6129484611666145821" }
    ) {
        id
        name
        course {
            name
        }
    }
}
```

```sh
go run github.com/99designs/gqlgen generate
```

N+1
Fix with dataloaders
