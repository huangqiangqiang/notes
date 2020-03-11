
# inner join 写法

用 gorm 的 inner join 写法如下：

```
db.Raw("select * from favorite_tbl, product_tbl where and favorite_tbl.product_id = product_tbl.id and favorite_tbl.user_id = '1'").Scan(&result)
```
这种是直接写 SQL 语句的，还有一种如下：

```
// Favorite 模型对应 favorite_tbl 表
result := []models.Favorite{}
db.Joins("join product_tbl on product_tbl.id = favorite_tbl.product_id and favorite_tbl.user_id = '1'").Find(&result)
```

# model 中的某个字段也是 model

如：Product 商品模型中有一个 brand 字段，brand也是一个model，那么需要先预加载 brand 表中的信息，再把 product 表和 `item_category_auth` 表 inner join 起来，后面再跟 where 过滤条件。

```
query := []models.Product{}
total := 0
db.
  // 预加载 Brand 表
  Preload("Brand").
  // inner join 后跟临时表的规则
  Joins("join item_category_auth on item_category_auth.item_id = item.id").
  // where 后跟实际过滤规则
  Where("item_category_auth.item_category_id = ?", cID).
  Offset((index - 1) * size).
  Limit(size).
  Find(&query).
  // 能获取的记录总数，不受 liimt 和 offset 限制
  Count(&total)
```

当然，model中也需要配置 foreignkey ，不然 product 中的 brand 字段怎么和 brand 表对应起来呢。

```
type Product struct {
	ID            int                     `json:"id" gorm:"column:id"`
	BrandID       int                     `json:"brand_id" gorm:"column:brand_id"`
	Brand         Brand                   `json:"brand" gorm:"foreignkey:BrandID"`
}

type Brand struct {
	ID        int       `json:"id" gorm:"column:id"`
}
```

指定了 Brand 字段的 foreignkey 为 BrandID ，BrandID 就会自动和 brand 表中的主键对应起来。