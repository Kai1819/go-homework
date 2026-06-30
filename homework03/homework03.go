package homework03

import (
	"fmt"
	"time"

	// "context"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// ### 题目1：模型定义
// - 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
//   - 要求 ：
//     - 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
//     - 编写Go代码，使用Gorm创建这些模型对应的数据库表。

type User struct {
	ID      uint64 `gorm:"primaryKey"`
	Name    string
	Posts    []Post `gorm:"foreignKey:UserID;references:ID"` // tag显式声明关联外键
	PostNum int64
}
type Post struct {
	ID      uint64 `gorm:"primaryKey"`
	UserID  uint64 `gorm:"index"` // 1对多，必须外键
	Title   string
	Content string
	Comments []Comment `gorm:"foreignKey:PostID;references:ID"` // tag显式声明关联外键
	Status  string
}
type Comment struct {
	ID      uint64 `gorm:"primaryKey"`
	PostID  uint64 `gorm:"index"` // 1对多，必须外键
	Content string
}

func GetSqlLiteDB() (db *gorm.DB, err error) {
	// 数据库
	db, err = gorm.Open(sqlite.Open("homework03.db"), &gorm.Config{
		// 日志配置
		Logger: logger.Default.LogMode(logger.Info),
		// 命名策略
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true, //false 允许 GORM 用默认的复数规则；表名后面有s
			NoLowerCase:   false,
			NameReplacer:  nil,
		},
	})
	if err != nil {
		// panic("连接数据库失败：" + err.Error())
		return nil, fmt.Errorf("连接数据库失败: %w", err)
	}

	// 连接池配置
	sqlDB, err := db.DB()
	if err != nil {
		//panic("打开通用数据库失败：" + err.Error())
		return nil, fmt.Errorf("打开通用数据库失败：: %w", err)
	}
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetMaxOpenConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	return
}

func Gorm01() {
	var db *gorm.DB
	var err error

	// 数据库
	db, err = GetSqlLiteDB()

	if err != nil {
		panic("获取数据库失败：" + err.Error())
	}

	// 关闭连接
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 自动建表
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		panic("自动建表和关联关系失败：" + err.Error())
	}

	// 关闭外键约束检查和删除数据
	db.Exec("PRAGMA foreign_keys = OFF")
	db.Exec("DELETE FROM User")
	db.Exec("DELETE FROM Post")
	db.Exec("DELETE FROM Comment")
	// 开启外键约束，需要关掉单独插入comments表
	db.Exec("PRAGMA foreign_keys = ON")

	// 插入数据
	// comments := []Comment{
	// 	{ID:1,Content: "中白友谊，坚如磐石！"},
	// 	{ID:2,Content: "互相尊重 和平发展"},
	// 	{ID:3,Content: "市场看好七月赢利多"},
	// }
	// err = db.Create(&comments).Error
	// if err != nil {
	// 	panic("Comment表插入数据失败：" + err.Error())
	// }
	user := User{
		ID:   1,
		Name: "深圳陈景川",
		Posts: []Post{
			{
				ID:      1,
				Title:   "卢卡申科低调空降北京！",
				Content: "中白是实打实全天候全面战略伙伴，风雨同舟、彼此撑腰。",
				Comments: []Comment{
					{ID: 1, Content: "中白友谊，坚如磐石！"},
					{ID: 2, Content: "互相尊重 和平发展"},
				},
				Status: "已评论",
			},
			{
				ID:      2,
				Title:   "7月市场主线已定！",
				Content: "昨天市场走的挺好的，领涨的就是半导体和医药。半导体设备和零部件已经说了好几周了，还没上车的朋友，就别上车了，别等你上车的时候，可能车就不开了。​半导体和医药大概率会是7月份的主线板块，但需要注意医药板块的标的不像半导体，内部肯定会分化。半导体设备主要还是要看7月上旬长鑫上市和业绩预告能不能超预期。目前已经进入加速上涨阶段，安心吃肉就行。",
				Comments: []Comment{
					{ID: 3, Content: "市场看好七月赢利多"},
				},
				Status: "已评论",
			},
		},
		PostNum: 0,
	}
	err = db.Session(&gorm.Session{FullSaveAssociations: true}).Create(&user).Error
	if err != nil {
		panic("User和Post表关联插入数据失败：" + err.Error())
	}

}

// ### 题目2：关联查询
// - 基于上述博客系统的模型定义。
//   - 要求 ：
//   - 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
//   - 编写Go代码，使用Gorm查询评论数量最多的文章信息。
func Gorm02() {
	var db *gorm.DB
	var err error

	// 数据库
	db, err = GetSqlLiteDB()

	if err != nil {
		panic("获取数据库失败：" + err.Error())
	}

	// 关闭连接
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	fmt.Printf("编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。")
	var user User
	// Preload 这个方法的参数必须是结构体里字段的名字，而不是关联的目标类型名。
	// Preload 的参数要用真实字段名（区分大小写），并且如果要预加载"二级关联"（比如通过 User → Posts → Comments），需要用点号连接路径：
	err = db.Preload("Posts").Preload("Posts.Comments").First(&user, "id = ?", 1).Error
	if err != nil {
		panic("查询失败：" + err.Error())
	}
	fmt.Println(user)
	for _,post := range user.Posts {
		fmt.Println(post)
		for _,comment  := range post.Comments {
			fmt.Println(comment)
		}
	}

	fmt.Printf("编写Go代码，使用Gorm查询评论数量最多的文章信息。")
	rows,err := db.Raw(`
		select p.id as id,count(c.id) as cnt
		from post p 
		left join comment c on p.id = c.post_id
		group by p.id
	`).Rows()
	if err != nil {
		panic("查询失败：" + err.Error())
	}
	defer rows.Close()
	for rows.Next() {
	    var id uint64
	    var cnt int64
	    err := rows.Scan(&id, &cnt)
	    if err != nil {
	        panic("Scan 失败：" + err.Error())
	    }
	    fmt.Println(id, cnt)
	}
    
	type PostWithCommentCount struct {
		Post
		CommentCount int64
	}
    var pwc PostWithCommentCount
	err = db.Raw(`
		select p.*, count(c.id) as comment_count
		from post p 
		left join comment c on p.id = c.post_id
		group by p.id
		order by comment_count desc
		limit 1
	`).Scan(&pwc).Error
	if err != nil {
		panic("查询失败：" + err.Error())
	}
	fmt.Println(pwc)


}

// - 继续使用博客系统的模型。
//   - 要求 ：
//     - 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
//     - 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func Gorm03() {
	var db *gorm.DB
	var err error

	// 数据库
	db, err = GetSqlLiteDB()

	if err != nil {
		panic("获取数据库失败：" + err.Error())
	}

	// 关闭连接
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	fmt.Println("为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。")
	post1 := Post{
		UserID:  1,
		Title:   "测试钩子函数",
		Content: "看看 PostNum 会不会自动+1",
		Status:  "草稿",
	}

	err = db.Create(&post1).Error
	if err != nil {
		panic("创建文章失败：" + err.Error())
	}

	var user User
	db.First(&user, "id = ?", 1)
	fmt.Println("当前 PostNum:", user.PostNum)

	fmt.Println("为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 无评论。")
	// 创建一个测试文章
	post2 := Post {
		ID: 10,
		UserID: 1,
		Title:   "测试文章-钩子测试",
		Content: "钩子测试",
		Comments: []Comment{
			{ID:11,Content: "评论1"},
			{ID:12,Content: "评论2"},
		},
		Status: "已评论",
	}
	err = db.Session(&gorm.Session{FullSaveAssociations: true}).Create(&post2).Error
	if err != nil {
		panic("极联新增Post和Comment表失败：" + err.Error())
	}
	// 先看一下某篇文章当前的评论数和状态
	var post Post
	db.Preload("Comments").First(&post, "id = ?", 10)
	fmt.Println("删除前 - 评论数:", len(post.Comments), "状态:", post.Status)

	// 把这篇文章下的所有评论都删掉
	// 只传条件，不传完整实例"的批量删除方式，GORM 在执行时走的是 DELETE FROM comments WHERE ... 这种 SQL，并不会先把每条记录的完整字段查出来再逐条触发钩子，所以钩子函数里能拿到的 c 实例字段大多是空的（除非你删除条件本身直接命中了某个字段，但这也不是稳定保证的行为，不同 GORM 版本可能有差异）
	// err = db.Where("post_id = ?", 10).Delete(&Comment{}).Error
	// 先查出完整的再删除，才会出发钩子
	err = DeleteCommentsByPostID(db,post.ID)
	if err != nil {
		panic("删除评论失败：" + err.Error())
	}

	// 再查一次，看看状态有没有自动更新
	db.First(&post, "id = ?", 10)
	fmt.Println("删除后 - 状态:", post.Status)
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Model(&User{}).
		Where("id = ?", p.UserID).
		UpdateColumn("post_num", gorm.Expr("post_num + ?", 1)).
		Error
}

/*
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	err = tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count).Error
	if err != nil {
		return fmt.Errorf("统计文章评论数量失败: %w", err)
	}

	if count == 0 {
		err = tx.Model(&Post{}).Where("id = ?", c.PostID).UpdateColumn("status", "无评论").Error
		if err != nil {
			return fmt.Errorf("更新文章状态失败: %w", err)
		}
	}

	return nil
}
*/

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println(c)
	var cnt int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&cnt).Error; err != nil {
		return err
	}
	if cnt == 0 {
		return tx.Model(&Post{}).Where("id = ?", c.PostID).UpdateColumn("status", "无评论").Error
	}
	return nil
}

func DeleteCommentsByPostID(db *gorm.DB, postID uint64) error {
	var comments []Comment
	if err := db.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		return fmt.Errorf("查询待删除评论失败: %w", err)
	}

	for _, c := range comments {
		if err := db.Delete(&c).Error; err != nil {
			return fmt.Errorf("删除评论失败: %w", err)
		}
	}

	return nil
}