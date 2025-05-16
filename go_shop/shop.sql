-- 创建数据库
DROP DATABASE IF EXISTS shop;
CREATE DATABASE shop CHARSET utf8;

-- 使用数据库
USE shop;

-- 创建数据表
-- 品牌表
CREATE TABLE brands (  
    brand_id INT AUTO_INCREMENT PRIMARY KEY,  
    english_name VARCHAR(100) NOT NULL,  
    chinese_name VARCHAR(100) NOT NULL,  
    grade char(2),  
    brand_story TEXT  
)charset utf8;

-- 插入数据
INSERT INTO brands (english_name, chinese_name, grade, brand_story) VALUES  
('Rolex', '劳力士', '高档', '劳力士（Rolex）是瑞士钟表业的经典品牌，以其卓越品质和创新精神著称。'),  
('Omega', '欧米茄', '中档', '欧米茄（Omega）是瑞士著名的钟表制造商，以其精湛的工艺和卓越的性能闻名。'),  
('Cartier', '卡地亚', '高档', '卡地亚（Cartier）是法国的珠宝和钟表制造商，以其独特的设计和精湛的工艺受到全球消费者的喜爱。'),  
('Patek Philippe', '百达翡丽', '奢华', '百达翡丽（Patek Philippe）是瑞士的顶级钟表制造商，以其复杂功能和精湛工艺著称。'),  
('Jaeger-LeCoultre', '积家', '高档', '积家（Jaeger-LeCoultre）是瑞士的钟表制造商，以其创新的机芯和卓越的品质闻名。'),  
('Audemars Piguet', '爱彼', '高档', '爱彼（Audemars Piguet）是瑞士的钟表制造商，以其精湛的制表工艺和独特的设计著称。'),  
('Breguet', '宝玑', '奢华', '宝玑（Breguet）是法国的钟表制造商，以其精湛的工艺和悠久的历史而闻名。'),  
('Chopard', '萧邦', '高档', '萧邦（Chopard）是瑞士的珠宝和钟表制造商，以其独特的设计和精湛的工艺受到全球消费者的喜爱。'),   
('longines', '浪琴', '中档', '浪琴表于1832年由奥古斯特·阿加西（Auguste Agassiz）在瑞士索伊米亚（Saint-Imier）创 立。他以自己的名字在瑞士开始经营钟表生意，随后他的侄儿欧内斯特·法兰西昂（Ernest Francillon）巧妙地把业务拓展为较具规模的钟表公司。1867年，浪琴表工厂正式在瑞士的Les Longines开设制表，并在同年以“Longines”名字注册，正式宣布诞生。'),  
('IWC Schaffhausen', '万国', '中档', '万国（IWC Schaffhausen）是瑞士的钟表制造商，以其精湛的制表工艺和独特的设计而著称。'),  
('Michel Herbelin', '法国赫柏林', '高档', '万国（IWC Schaffhausen）是瑞士的钟表制造商，以其精湛的制表工艺和独特的设计而著称。'),  
('Frank Muller', '法兰克·穆勒', '奢华', '1985年，创始人米歇尔·赫柏林退休，儿子让-克劳德·赫柏林（Jean-Claude Herbelin）继承 独立制表家族企业，担任总裁。1987年，首款Newport腕表系列诞生，该系列设计优雅且具魄力，有多个不同款式，如女装表、男装表、 石英表、机械表等。');

-- 系列表
CREATE TABLE series (  
    series_id INT AUTO_INCREMENT PRIMARY KEY,  
    brand_id INT,  
    series_name VARCHAR(100) NOT NULL,  
    FOREIGN KEY (brand_id) REFERENCES brands(brand_id)  
);

-- 插入系列数据
INSERT INTO series (brand_id, series_name) VALUES  
(1, 'Submariner'),  
(1, 'Daytona'),  
(2, 'Seamaster'),  
(2, 'Constellation'),  
(3, 'Panthère'),  
(3, 'Santos'),  
(4, 'Calatrava'),  
(4, 'Nautilus'),  
(5, 'Master Control'), 
(12,'Passion'), 
(12,'Vilentines'), 
(5, 'Reverso');

-- 商品表
CREATE TABLE products (  
    product_id INT AUTO_INCREMENT PRIMARY KEY,  
    series_id INT COMMENT '系列',  
    case_material VARCHAR(10) COMMENT '表壳',  
    case_back VARCHAR(10) COMMENT '表底',  
    strap_color VARCHAR(10) COMMENT '表带颜色',  
    functions TEXT COMMENT '功能',  
    style VARCHAR(10) COMMENT '款式',  
    size VARCHAR(20) COMMENT '尺寸',  
    watch_glass VARCHAR(20) COMMENT '表镜',  
    watch_buckle VARCHAR(10) COMMENT '表扣',  
    launch_year YEAR COMMENT '推出年份',  
    movement VARCHAR(10) COMMENT '机芯',  
    thickness DECIMAL(5,2) COMMENT '厚度',  
    dial VARCHAR(10) COMMENT '表盘',  
    water_resistance VARCHAR(50) COMMENT '防水',  
    price DECIMAL(10,2) COMMENT '价格',  
    is_new BOOLEAN DEFAULT TRUE COMMENT '新品',  
    is_special_offer BOOLEAN DEFAULT FALSE COMMENT '特价', 
    special_price DECIMAL(10,2) DEFAULT 0 COMMENT '特价价格',
    image VARCHAR(100) COMMENT '图片名字',
    gender varchar(10) COMMENT '男女款式',
    ordered_num INT UNSIGNED DEFAULT 0 COMMENT '已售数量',
    FOREIGN KEY (series_id) REFERENCES series(series_id)  
);

-- 插入商品数据
INSERT INTO products (series_id,ordered_num, gender,image,case_material, case_back, strap_color, functions, style, size, watch_glass, watch_buckle, launch_year, movement, thickness, dial, water_resistance, price, is_new, is_special_offer) VALUES  
(3, 10,'男','3675_37867.jpg','精钢', '密底', 'drop银色', '石英', '简约', '28mm', '矿物玻璃', '针扣', 2022, '石英', 6.00, '银色', '30米', 3999.00, TRUE, FALSE),  
(3, 200,'女','3195_1BS_89881.jpg','黄金与精钢', '透底', '玫瑰金', '石英, 日期显示', '优雅', '33mm', '蓝宝石水晶玻璃', '折叠扣', 2021, '石英', 7.00, '白色', '50米', 6999.00, TRUE, FALSE),  
(4, 10,'女','E-221_23889.jpg','不锈钢', '密底', '黑色', '自动上链, 计时器', '运动', '44mm', '蓝宝石水晶玻璃', '折叠扣', 2020, '机械', 13.50, '黑色', '200米', 9999.00, TRUE, FALSE),  
(4, 30,'女','7718_63938.jpg','18K白金', '透底', '白色', '自动上链, 计时器, 日期显示', '豪华', '38mm', '蓝宝石水晶玻璃', '隐藏式扣', 2019, '机械', 10.00, '蓝色', '100米', 29999.00, TRUE, FALSE),  
(5, 10,'女','12443_S01_97254.jpg','精钢', '密底', '棕色', '自动上链, 日期显示, 动力储备显示', '经典', '42mm', '蓝宝石水晶玻璃', '折叠扣', 2023, '机械', 11.00, '银色', '50米', 8999.00, TRUE, FALSE),  
(5, 1,'男','16155550_16813.jpg','玫瑰金', '透底', '红色', '自动上链, 计时器, 日期显示, 月相', '复杂功能', '40mm', '蓝宝石水晶玻璃', '针扣', 2022, '机械', 12.50, '红色', '30米', 49999.00, TRUE, FALSE),  
(6, 10,'男','7718_63938.jpg','白金', '密底', '黑色', '陀飞轮, 日期显示', '高端复杂', '47mm', '蓝宝石水晶玻璃', '折叠扣', 2021, '机械', 14.00, '银色', '50米', 199999.00, TRUE, TRUE),  
(6, 10,'女','7718_63938.jpg','黄金', '透底', '金色', '三问报时, 日期显示', '艺术珍品', '45mm', '蓝宝石水晶玻璃', '隐藏式扣', 2020, '机械', 15.00, '金色', '30米', 399999.00, TRUE, FALSE),  
(7, 52,'男','3317_92065.jpg','精钢', '密底', '蓝色', '石英, 日期显示', '简约时尚', '36mm', '矿物玻璃', '针扣', 2023, '石英', 5.50, '蓝色', '50米', 2999.00, TRUE, FALSE),  
(7, 10,'男','212_30_41_20_03_001_79603.jpg','玫瑰金与精钢', '透底', '粉色', '石英, 日期显示, 闹钟', '女性专属', '28mm', '蓝宝石水晶玻璃', '折叠扣', 2022, '石英', 6.50, '粉色', '30米', 4999.00, TRUE, FALSE),  
(9, 10,'男','9740A-AG-M9740_43172.jpg','18K黄金', '透底', '绿色', '自动上链, 计时器, 日期显示, 年历', '豪华复杂', '43mm', '蓝宝石水晶玻璃', '折叠扣', 2018, '机械', 13.00, '绿色', '100米', 99999.00, TRUE, FALSE),  
(10,10,'男','M1-37-34-LB_62213.jpg', '铂金', '密底', '紫色', '陀飞轮, 三问报时, 日期显示', '顶级艺术', '48mm', '蓝宝石水晶玻璃', '隐藏式扣', 2017, '机械', 16.00, '紫色', '50米', 599999.00, TRUE, TRUE),
(3, 80,'男','3675_37867.jpg','精钢', '密底', '银色', '石英', '简约', '28mm', '矿物玻璃', '针扣', 2022, '石英', 6.00, '银色', '30米', 3999.00, TRUE, FALSE),  
(3, 10,'男','3195_1BS_89881.jpg','黄金与精钢', '透底', '玫瑰金', '石英, 日期显示', '优雅', '33mm', '蓝宝石水晶玻璃', '折叠扣', 2021, '石英', 7.00, '白色', '50米', 6999.00, TRUE, FALSE),  
(4, 43,'男','E-221_23889.jpg','不锈钢', '密底', '黑色', '自动上链, 计时器', '运动', '44mm', '蓝宝石水晶玻璃', '折叠扣', 2020, '机械', 13.50, '黑色', '200米', 9999.00, TRUE, FALSE),  
(11, 10,'女','7718_63938.jpg','18K白金', '透底', '白色', '自动上链, 计时器, 日期显示', '豪华', '38mm', '蓝宝石水晶玻璃', '隐藏式扣', 2019, '机械', 10.00, '蓝色', '100米', 29999.00, TRUE, FALSE),  
(5, 8,'女','12443_S01_97254.jpg','精钢', '密底', '棕色', '自动上链, 日期显示, 动力储备显示', '经典', '42mm', '蓝宝石水晶玻璃', '折叠扣', 2023, '机械', 11.00, '银色', '50米', 8999.00, TRUE, FALSE),  
(5, 12,'女','16155550_16813.jpg','玫瑰金', '透底', '红色', '自动上链, 计时器, 日期显示, 月相', '复杂功能', '40mm', '蓝宝石水晶玻璃', '针扣', 2022, '机械', 12.50, '红色', '30米', 49999.00, TRUE, FALSE),  
(6, 10,'女','7718_63938.jpg','白金', '密底', '黑色', '陀飞轮, 日期显示', '高端复杂', '47mm', '蓝宝石水晶玻璃', '折叠扣', 2021, '机械', 14.00, '银色', '50米', 199999.00, TRUE, TRUE),  
(12, 10,'女','7718_63938.jpg','黄金', '透底', '金色', '三问报时, 日期显示', '艺术珍品', '45mm', '蓝宝石水晶玻璃', '隐藏式扣', 2020, '机械', 15.00, '金色', '30米', 399999.00, TRUE, FALSE),  
(7, 10,'男','3317_92065.jpg','精钢', '密底', '蓝色', '石英, 日期显示', '简约时尚', '36mm', '矿物玻璃', '针扣', 2023, '石英', 5.50, '蓝色', '50米', 2999.00, TRUE, FALSE),  
(7, 10,'男','212_30_41_20_03_001_79603.jpg','玫瑰金与精钢', '透底', '粉色', '石英, 日期显示, 闹钟', '女性专属', '28mm', '蓝宝石水晶玻璃', '折叠扣', 2022, '石英', 6.50, '粉色', '30米', 4999.00, TRUE, FALSE),  
(9, 0,'男','9740A-AG-M9740_43172.jpg','18K黄金', '透底', '绿色', '自动上链, 计时器, 日期显示, 年历', '豪华复杂', '43mm', '蓝宝石水晶玻璃', '折叠扣', 2018, '机械', 13.00, '绿色', '100米', 99999.00, TRUE, FALSE),  
(10,201,'女','M1-37-34-LB_62213.jpg', '铂金', '密底', '紫色', '陀飞轮, 三问报时, 日期显示', '顶级艺术', '48mm', '蓝宝石水晶玻璃', '隐藏式扣', 2017, '机械', 16.00, '紫色', '50米', 599999.00, TRUE, TRUE);  

-- 用户表
CREATE TABLE users (  
    user_id INT AUTO_INCREMENT PRIMARY KEY,  
    username VARCHAR(30) NOT NULL UNIQUE,  
    password VARCHAR(128) NOT NULL,  
    email VARCHAR(50) UNIQUE,  
    phone_number char(11) UNIQUE  
);

-- 插入一条用户记录
INSERT INTO users VALUES(null,"admin",md5("123456"),"12345@qq.com","13212345678");

-- 购物车表

CREATE TABLE carts (  
    cart_id INT AUTO_INCREMENT PRIMARY KEY,  
    user_id INT,  
    product_id INT,  
    quantity INT DEFAULT 1,  
    FOREIGN KEY (user_id) REFERENCES users(user_id),  
    FOREIGN KEY (product_id) REFERENCES products(product_id)  
);


-- 订单表

CREATE TABLE orders (  
    order_id INT AUTO_INCREMENT PRIMARY KEY,  
    user_id INT,  
    order_date DATETIME DEFAULT CURRENT_TIMESTAMP,  
    total_price DECIMAL(15,2),  
    FOREIGN KEY (user_id) REFERENCES users(user_id)  
);

-- 订单商品关联表
CREATE TABLE order_products (  
    order_product_id INT AUTO_INCREMENT PRIMARY KEY,  
    order_id INT,  
    product_id INT,  
    quantity INT,  
    price DECIMAL(15,2),  
    FOREIGN KEY (order_id) REFERENCES orders(order_id),  
    FOREIGN KEY (product_id) REFERENCES products(product_id)  
);

-- 获取首页相关数据
-- 首页需要显示的数据：商品ID、品牌名字、系列名字、商品名字、商品价格、商品销量、商品图片、商品类别、是否特价
select p.product_id,p.image,p.price,p.is_special_offer,p.special_price,p.gender,p.style,p.ordered_num,s.series_name,b.chinese_name,b.grade from products p left join series s on p.series_id = s.series_id left join brands b on s.brand_id = b.brand_id;

-- 修改数据的价格(rand() 是0-1之间的随机小数)
update products set special_price = price - floor(rand() * 10000) where is_special_offer = 1;

-- 获取所有特价商品数据
select p.product_id,p.style,p.price,p.is_special_offer,p.special_price,p.image,p.ordered_num,s.series_name,b.chinese_name,b.english_name from products p left join series s on p.series_id = s.series_id left join brands b on s.brand_id = b.brand_id where p.is_special_offer = 1;


-- 更新数据：增加新品数据
update products set is_new = round(rand());

-- 获取商品的详情
select from products p left join series s on p.series_id = s.series_id left join brands b on s.brand_id = b.brand_id where p.product_id = 

-- 获取指定商品的热销
select p.product_id,p.image,p.price,p.is_special_offer,p.special_price,p.gender,p.style,p.ordered_num,s.series_name,b.chinese_name,b.grade from products p left join series s on p.series_id = s.series_id left join brands b on s.brand_id = b.brand_id = ? order by p.ordered_num desc limait 3

