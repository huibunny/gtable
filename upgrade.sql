-- 2022-09-16
CREATE TABLE public.t_kv (
	k varchar(32) NULL,
	v jsonb NULL,
	CONSTRAINT t_kv_pk PRIMARY KEY (k)
);

-- 2022-09-17
CREATE TABLE public.t_color (
	id uuid NOT NULL default gen_random_uuid(),
	color varchar(16) NULL,
	color_cn varchar(16) NULL,
	color_desc varchar(256) NULL,
	CONSTRAINT t_color_pk PRIMARY KEY (id)
);

INSERT INTO public.t_color
(color, color_cn, color_desc)
VALUES
('black', '黑色', '黑色 - 生活相关，黑白的生活是简单的。'),
('red', '红色', '红色 - 健康相关，红白的健康值得注意。'),
('green', '绿色', '绿色 - 隐私相关，绿了就绿了，那是你的隐私。'),
('blue', '蓝色', '蓝色 - 学习相关，畅游知识的海洋。');


