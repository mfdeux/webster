package store

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/olivere/elastic/v7"
)

type ElasticService struct {
	Host   string
	Port   int
	client *elastic.Client
}

func MakeDateIndex(name string) string {
	currentTime := time.Now()
	return fmt.Sprintf("%s-%s", name, currentTime.Format("01-02-2006"))
}

func (e *ElasticService) IndexDocument(index string, id string, doc interface{}) (bool, error) {
	ctx := context.Background()
	_, err := e.client.Index().
		Index(index).
		Id(id).
		BodyJson(doc).
		Refresh("true").
		Do(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (e *ElasticService) DeleteDocument(index string, id string) (bool, error) {
	ctx := context.Background()
	_, err := e.client.Delete().
		Index(index).
		Id(id).
		Refresh("true").
		Do(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (e *ElasticService) Ping() error {
	ctx := context.Background()
	_, _, err := e.client.Ping("http://127.0.0.1:9200").Do(ctx)
	return err
}

func (e *ElasticService) Flush(index string) error {
	ctx := context.Background()
	_, err := e.client.Flush().Index(index).Do(ctx)
	return err
}

func (e *ElasticService) EnsureDocIndex(name string) (bool, error) {
	ctx := context.Background()
	exists, err := e.client.IndexExists(name).Do(ctx)
	if err != nil {
		return false, err
	}
	if !exists {
		// createIndex, err := e.client.CreateIndex(name).Body(testMapping).Do(ctx)
		createIndex, err := e.client.CreateIndex(name).Do(ctx)
		if err != nil {
			return false, err
		}
		if !createIndex.Acknowledged {
			return false, nil
		}
	}
	return true, nil
}

func (e *ElasticService) EnsureDocIndexWithBody(name string, body string) (bool, error) {
	ctx := context.Background()
	exists, err := e.client.IndexExists(name).Do(ctx)
	if err != nil {
		return false, err
	}
	if !exists {
		createIndex, err := e.client.CreateIndex(name).Body(body).Do(ctx)
		if err != nil {
			return false, err
		}
		if !createIndex.Acknowledged {
			return false, nil
		}
	}
	return true, nil
}

func (e *ElasticService) EnsureDocIndexWithBodyFile(name string, filepath string) (bool, error) {
	bodyBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return false, err
	}
	return e.EnsureDocIndexWithBody(name, string(bodyBytes))
}

func (e *ElasticService) DeleteDocIndex(name string) (bool, error) {
	ctx := context.Background()
	deleteIndex, err := e.client.DeleteIndex(name).Do(ctx)
	if err != nil {
		return false, err
	}
	if !deleteIndex.Acknowledged {
		return false, nil
	}
	return true, nil
}

func (e *ElasticService) Client() *elastic.Client {
	return e.client
}

func NewElasticService() (*ElasticService, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	return &ElasticService{
		client: client,
	}, err
}

func (e *ElasticService) EnsurePercolateIndexWithBody(name string, body string) (bool, error) {
	return e.EnsureDocIndexWithBody(name, body)
}

func (e *ElasticService) EnsurePercolateIndexWithBodyFile(name string, filepath string) (bool, error) {
	return e.EnsureDocIndexWithBodyFile(name, filepath)
}

func (e *ElasticService) DeletePercolateIndex(name string) (bool, error) {
	return e.DeleteDocIndex(name)
}

func (e *ElasticService) PercolateDocument(index string, doc interface{}) (*elastic.SearchHits, error) {
	ctx := context.Background()
	pq := elastic.NewPercolatorQuery().
		Field("query").
		Document(doc)

	res, err := e.client.Search(index).Query(pq).Do(ctx)
	return res.Hits, err
}
func (e *ElasticService) PercolateIndexedDocument(index string, id string) (*elastic.SearchHits, error) {
	ctx := context.Background()
	pq := elastic.NewPercolatorQuery().
		Field("query").IndexedDocumentId(id)

	res, err := e.client.Search(index).Query(pq).Do(ctx)
	return res.Hits, err
}

// termQuery := elastic.NewTermQuery("user", "olivere")
// 	searchResult, err := client.Search().
// 		Index("twitter"). // search in index "twitter"
// 		Query(termQuery). // specify the query
// 		Do(ctx)           // execute
// 	if err != nil {
// 		fmt.Println("searching")
// 		panic(err)
// 	}

// Update Mapping
// Query
