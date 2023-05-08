/**
* Created by GoLand
* @file base_model.go
* @version: 1.0.0
* @author 李锦 <Lijin@cavemanstudio.net>
* @date 2022/1/27 2:57 下午
* @desc base_model.go
 */

package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-xmodule/utils/dirver"
	"github.com/go-xmodule/utils/global"
	"github.com/go-xmodule/utils/utils/xerror"
	"github.com/go-xmodule/utils/utils/xlog"
	"gorm.io/gorm"
	"math"
	"time"
)

type Where any

// Database 基础模型
type Database struct {
	db     *gorm.DB
	Result any
}

// DBHandler 数据库操作句柄
var DBHandler *Database

type ModelAction interface {
	DataId() int
	TableName() string
}

// CommonModel 数据库类型基类
type CommonModel struct {
	Id        int       `gorm:"column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// InitializeMysql 初始化数据库连接
func InitializeMysql(config dirver.LinkParams) *gorm.DB {
	db, err := dirver.InitializeDB(config)
	xerror.PanicErr(err, global.InitMysqlErr.String())
	DBHandler = NewDatabase(db)
	return db
}

// NewDatabase 获取新的模型
func NewDatabase(db *gorm.DB) *Database {
	database := new(Database)
	database.db = db
	database.Result = &[]map[string]any{}
	return database
}

// SetModel 设置数据结果model
func (d *Database) SetModel(executeModel any) *Database {
	d.Result = executeModel
	return d
}

// SetResult 设置数据结果model
func (d *Database) SetResult(resultModel any) *Database {
	d.Result = resultModel
	return d
}

// AutoMigrate 生成表结构
func (d *Database) AutoMigrate() error {
	err := d.db.AutoMigrate(d.Result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteByWhere 删除数据
func (d *Database) DeleteByWhere(where any) error {
	err := d.db.Where(where).Delete(d.Result).Error
	if err != nil {
		xlog.Logger.WithField("err", err).Error(global.DataDeleteErr.String())
	}
	return nil
}

// First 根据条件获取一条数据
func (d *Database) First(where Where, not ...Where) error {
	if len(not) > 0 {
		return d.db.Where(where).Not(not[0]).First(d.Result).Error
	} else {
		return d.db.Where(where).First(d.Result).Error
	}
}

func (d *Database) Begin() *gorm.DB {
	return d.db.Begin()
}

// Exist 检查数据是否存在
func (d *Database) Exist(where Where) (bool, error) {
	var count int64
	err := d.db.Model(d.Result).Where(where).Count(&count).Error
	if err != nil {
		xlog.Logger.WithField("err", err).Error(global.DbErr.String())
	}
	return count > 0, nil
}

// ExecuteSql 执行sql
func (d *Database) ExecuteSql(sql string, logMod ...bool) error {
	return d.db.Raw(sql).Scan(d.Result).Error
}

// Find 根据id获取一条数据
func (d *Database) Find(id any) error {
	return d.db.Where("id=?", id).First(d.Result).Error
}

// Get 根据where条件查询数据
func (d *Database) Get(where ...any) error {
	if len(where) > 0 {
		return d.db.Where(where[0]).Find(d.Result).Error
	} else {
		return d.db.Find(d.Result).Error
	}
}

// GetDb 获取当前数据库连接
func (d *Database) GetDb() *gorm.DB {
	return d.db
}

// Create 创建数据
func (d *Database) Create(model any) error {
	return d.db.Create(model).Error
}

// Save 保存修改的数据
func (d *Database) Save(model any) error {
	return d.db.Save(model).Error
}

// Delete 删除数据
func (d *Database) Delete(model any) error {
	return d.db.Delete(model).Error
}

// 获取表注释

func (d *Database) GetTableComment(database string, table string) (string, error) {
	var tableName string
	sqlFormat := "select table_comment from  information_schema.TABLES where TABLE_SCHEMA ='%s' and TABLE_NAME='%s'"
	sql := fmt.Sprintf(sqlFormat, database, table)
	err := d.SetResult(&tableName).ExecuteSql(sql)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

type Field struct {
	ColumnName    string `json:"column_name"`
	DataType      string `json:"data_type"`
	ColumnComment string `json:"column_comment"`
}

func (d *Database) GetTableFieldComment(database string, table string) ([]Field, error) {
	sqlFormat := "select column_name,data_type,column_comment from information_schema.COLUMNS where TABLE_SCHEMA ='%s' and TABLE_NAME='%s'"
	sql := fmt.Sprintf(sqlFormat, database, table)
	var result []Field
	err := d.SetResult(&result).ExecuteSql(sql)
	return result, err
}

// PaginationQuery 分页查询
type PaginationQuery struct {
	PageSize int
	PageNum  int
	// OrderBy 小写的字段名称
	OrderBy string
	// Order 默认是'desc', 可选的: 'desc', 'asc'
	Order string
}
type PageInfo struct {
	PrevPage   int
	CurrPage   int
	NextPage   int
	Pages      []int
	TotalCount int
	TotalPage  int
	Offset     int
	EndOffset  int
}
type PageData struct {
	PageInfo PageInfo
	DataList any
}

// GetByPage 分页查询数据
func (d *Database) GetByPage(pagination PaginationQuery, where any) PageData {
	var count int64
	d.db.Model(d.Result).Where(where).Count(&count)
	offset := pagination.PageSize * (pagination.PageNum - 1)
	d.db.Where(where).Order(fmt.Sprintf("%s %s ", pagination.OrderBy, pagination.Order)).Offset(offset).Limit(pagination.PageSize).Find(d.Result)
	paginator := d.paginator(pagination.PageNum, pagination.PageSize, int64(count))
	paginator.Offset = offset
	paginator.EndOffset = pagination.PageSize + offset
	data, _ := json.Marshal(d.Result)
	mapRes := make([]map[string]any, 0)
	_ = json.Unmarshal(data, &mapRes)
	return PageData{
		PageInfo: paginator,
		DataList: mapRes,
	}
}
func (d *Database) SetMode() error {
	sql := "set sql_mode ='STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';"
	return d.ExecuteSql(sql)
}

func (d *Database) paginator(currentPage, prePage int, totalCount int64) PageInfo {
	var prevPage int // 前一页地址
	var nextPage int // 后一页地址
	// 根据totalCount总数，和prePage每页数量 生成分页总数
	totalPage := int(math.Ceil(float64(totalCount) / float64(prePage))) // page总数
	if currentPage > totalPage {
		currentPage = totalPage
	}
	if currentPage <= 0 {
		currentPage = 1
	}
	var pages []int
	switch {
	case currentPage >= totalPage-5 && totalPage > 5: // 最后5页
		start := totalPage - 5 + 1
		prevPage = currentPage - 1
		nextPage = int(math.Min(float64(totalPage), float64(currentPage+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case currentPage >= 3 && totalPage > 5:
		start := currentPage - 3 + 1
		pages = make([]int, 5)
		prevPage = currentPage - 3
		for i, _ := range pages {
			pages[i] = start + i
		}
		prevPage = currentPage - 1
		nextPage = currentPage + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalPage))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
		prevPage = int(math.Max(float64(1), float64(currentPage-1)))
		nextPage = currentPage + 1
	}
	return PageInfo{
		PrevPage:   prevPage,
		NextPage:   nextPage,
		TotalPage:  totalPage,
		CurrPage:   currentPage,
		Pages:      pages,
		TotalCount: int(totalCount),
	}
}
