package pac_service

import (
	"fmt"
	"strings"
	"time"

	"github.com/haodiaodemingzi/cloudfeet/common/logging"
	"github.com/haodiaodemingzi/cloudfeet/models"
)

var logger = logging.GetLogger()

type Pac struct {
	ID         int       `db:"id"`
	Domain     string    `db:"domain"`
	Region     string    `db:"region"`
	Status     string    `db:"status"`
	Source     string    `db:"source"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

// SavePacDomain ...
func SavePacDomain(source string, domains string) error{
	var pac = &models.Pac{}
	var domainList = strings.Split(domains, ",")
	var data []map[string]interface{}
	for _, item := range domainList {
		data = append(data, map[string]interface{}{
			"domain": item,
			"status": 0,
			"source": source,
		})
	}
	_, err := pac.AddDomains(data)
	return err
}

// SavePacDomain ...
func GetDomains(cond map[string]interface{}) ([]string, error) {
	var pac = &models.Pac{}
	pacList, err := pac.Query(cond)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var domainList []string
	for i := 0; i < len(pacList); i++ {
		domainList = append(domainList, pacList[i].Domain)
	}

	return domainList, nil
}

// 刷新域名检测
func RefreshCheckedDomain(data *map[string]interface{}) error {
	var pac = &models.Pac{}
	return pac.UpdateCheckedDomain(data)
}
