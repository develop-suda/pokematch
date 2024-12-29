
### pokemons
#### 概要
- ポケモン一覧テーブル
- 名前、身長、体重情報を格納

| No  | PK  | FK  | カラム名   | 備考 | データ型     | NOT NULL | 列制約 |
| --- | --- | --- | ---------- | ---- | ------------ | -------- | ------ |
| 1   | ⚪︎   |     | id         |      | bigint       | ⚪︎        |        |
| 2   |     |     | name       |      | text         | ⚪︎        |        |
| 3   |     |     | height     |      | numeric(3,1) | ⚪︎        |        |
| 4   |     |     | weight     |      | numeric(4,1) | ⚪︎        |        |
| 5   |     |     | created_at |      | timestamp    | ⚪︎        |        |
| 6   |     |     | updated_at |      | timestamp    | ⚪︎        |        |
| 7   |     |     | deleted_at |      | timestamp    | ⚪︎        |        |
