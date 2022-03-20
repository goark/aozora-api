# [aozora-api] -- APIs for Aozora-bunko RESTful Service by Golang

[![check vulns](https://github.com/goark/aozora-api/workflows/vulns/badge.svg)](https://github.com/goark/aozora-api/actions)
[![lint status](https://github.com/goark/aozora-api/workflows/lint/badge.svg)](https://github.com/goark/aozora-api/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/goark/aozora-api/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/goark/aozora-api.svg)](https://github.com/goark/aozora-api/releases/latest)

This package is required Go 1.16 or later.

**Migrated repository to [github.com/goark/aozora-api][aozora-api]**

## Usage of package

### Import Package

```go
import "github.com/goark/aozora-api"
```

### Search for Aozora-bunko Books Data

```go
books, err := aozora.DefaultClient().SearchBooks(
    aozora.WithBookTitle("/天に積む宝/"),
    aozora.WithBookAuthor("富田倫生"),
)
```

### Lookup Aozora-bunko Book Data

```go
book, err := aozora.DefaultClient().LookupBook(59489)
```

### Search for Aozora-bunko Persons Data

```go
persons, err := aozora.DefaultClient().SearchPersons(
    aozora.WithPersonName("富田倫生"),
)
```

### Lookup Aozora-bunko Person Data

```go
person, err := aozora.DefaultClient().LookupPerson(55)
```

### Search for Aozora-bunko Workers Data

```go
workers, err := aozora.DefaultClient().SearchWorkers(
    aozora.WithWorkerName("雪森"),
)
```

### Lookup Aozora-bunko Worker Data

```go
worker, err := aozora.DefaultClient().LookupWorker(845)
```

### Lookup Ranking data of Aozora-bunko

```go
tm, err := time.Parse("2006-01", "2019-01")
ranking, err := aozora.DefaultClient().Ranking(tm)
```

## Entities for Aozora-bunko

### Book type

```go
//Author is entity class of author and translator info.
type Author struct {
    PersonID  int    `json:"person_id"`
    LastName  string `json:"last_name"`
    FirstName string `json:"first_name"`
}

//Book is entity class of book info.
type Book struct {
    BookID                      int      `json:"book_id"`
    Title                       string   `json:"title"`
    TitleYomi                   string   `json:"title_yomi"`
    TitleSort                   string   `json:"title_sort"`
    Subtitle                    string   `json:"subtitle"`
    SubtitleYomi                string   `json:"subtitle_yomi"`
    OriginalTitle               string   `json:"original_title"`
    FirstAppearance             string   `json:"first_appearance"`
    NDCCode                     string   `json:"ndc_code"`
    FontKanaType                string   `json:"font_kana_type"`
    Copyright                   bool     `json:"copyright"`
    ReleaseDate                 Date     `json:"release_date"`
    LastModified                Date     `json:"last_modified"`
    CardURL                     string   `json:"card_url"`
    BaseBook1                   string   `json:"base_book_1"`
    BaseBookPublisher1          string   `json:"base_book_1_publisher"`
    BaseBookFirstEdition1       string   `json:"base_book_1_1st_edition"`
    BaseBookEditionInput1       string   `json:"base_book_1_edition_input"`
    BaseBookEditionProofing1    string   `json:"base_book_1_edition_proofing"`
    BaseBookParent1             string   `json:"base_book_1_parent"`
    BaseBookParentPublisher1    string   `json:"base_book_1_parent_publisher"`
    BaseBookParentFirstEdition1 string   `json:"base_book_1_parent_1st_edition"`
    BaseBook2                   string   `json:"base_book_2"`
    BaseBookPublisher2          string   `json:"base_book_2_publisher"`
    BaseBookFirstEdition2       string   `json:"base_book_2_1st_edition"`
    BaseBookEditionInput2       string   `json:"base_book_2_edition_input"`
    BaseBookEditionProofing2    string   `json:"base_book_2_edition_proofing"`
    BaseBookParent2             string   `json:"base_book_2_parent"`
    BaseBookParentPublisher2    string   `json:"base_book_2_parent_publisher"`
    BaseBookParentFirstEdition2 string   `json:"base_book_2_parent_1st_edition"`
    Input                       string   `json:"input"`
    Proofing                    string   `json:"proofing"`
    TextURL                     string   `json:"text_url"`
    TextLastModified            Date     `json:"text_last_modified"`
    TextEncoding                string   `json:"text_encoding"`
    TextCharset                 string   `json:"text_charset"`
    TextUpdated                 int      `json:"text_updated"`
    HTMLURL                     string   `json:"html_url"`
    HTMLLastModified            Date     `json:"html_last_modified"`
    HTMLEncoding                string   `json:"html_encoding"`
    HTMLCharset                 string   `json:"html_charset"`
    HTMLUpdated                 int      `json:"html_updated"`
    Translators                 []Author `json:"translators"`
    Authors                     []Author `json:"authors"`
}
```

### Person type

```go
//Person is entity class of person info.
type Person struct {
    PersonID        int    `json:"person_id"`
    LastName        string `json:"last_name"`
    FirstName       string `json:"first_name"`
    LastNameYomi    string `json:"last_name_yomi"`
    FirstNameYomi   string `json:"first_name_yomi"`
    LastNameSort    string `json:"last_name_sort"`
    FirstNameSort   string `json:"first_name_sort"`
    LastNameRoman   string `json:"last_name_roman"`
    FirstNameRoman  string `json:"first_name_roman"`
    DateOfBirth     Date   `json:"date_of_birth"`
    DateOfDeath     Date   `json:"date_of_death"`
    AuthorCopyright bool   `json:"author_copyright"`
}
```

### Worker type

```go
//Worker is entity class of worker info.
type Worker struct {
    WorkerID int    `json:"id"`
    Name     string `json:"name"`
}
```

### Ranking type

```go
//Ranking is entity class of ranking info.
type Ranking []struct {
    BookID  int      `json:"book_id"`
    Access  int      `json:"access"`
    Title   string   `json:"title"`
    Authors []string `json:"authors"`
}
```

## Command Line Interface (Sample Code)

### Install and Build

```
$ go get github.com/goark/aozora-api/cli/aozora-bunko
```

### Search for Aozora-bunko Books Data

```
$ aozora-bunko search books -t "/天に積む宝/" -a "富田倫生"
```

### Lookup Aozora-bunko Book Data

```
$ aozora-bunko lookup book 59489
```

### Search for Aozora-bunko Persons Data

```
$ aozora-bunko search persons -n "富田倫生"
```

### Lookup Aozora-bunko Person Data

```
$ aozora-bunko lookup person 55
```

### Search for Aozora-bunko Workers Data

```
$ aozora-bunko search workers -n "雪森"
```

### Lookup Aozora-bunko Worker Data

```
$ aozora-bunko lookup worker 845
```

### Lookup Ranking data of Aozora-bunko

```
$ aozora-bunko ranking 2019-01
```

## Reference

- [aozorahack/pubserver2: Pubserver](https://github.com/aozorahack/pubserver2)
- [aozorahack/aozora-cli](https://github.com/aozorahack/aozora-cli) : CLI by Python

[aozora-api]: https://github.com/goark/aozora-api "goark/aozora-api: APIs for Aozora-bunko RESTful Service by Golang"
