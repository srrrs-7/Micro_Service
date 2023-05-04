le# directory structure

### cmd

cmd ディレクトリは、サービスの実行可能なエントリポイントを定義するためのディレクトリです。各サービスのエントリポイントをこのディレクトリに配置します。各エントリポイントは、サービスの初期化処理を実行するために使用されます。また、サービスの起動前に必要な環境変数の読み込みや、コマンドライン引数の解析などもこのディレクトリで実装します。

### app

internal ディレクトリは、外部からアクセスされない内部の実装のためのディレクトリです。ここに、ドメイン層、リポジトリ層、ユースケース層、デリバリー層、エンティティ層などのパッケージを作成します。各層が担う責務を明確にすることで、機能の独立性、テストのやりやすさ、保守性を高めることができます。

#### entity 層

エンティティ層は、ドメイン層で使用されるデータ構造を定義する層です。各エンティティは、ユニークな ID を持ちます。

#### repository 層

リポジトリ層は、データベースやストレージなどの永続化層を担当します。この層では、永続化されたデータに対する処理を実装します。例えば、ユーザーの作成・更新・削除などがあります。

#### domain 層

ドメイン層は、ビジネスロジックを実装するための層です。例えば、ユーザー情報のバリデーション、検証、処理などが担当する層です。この層では、必要なデータ構造やインターフェースを定義します。

#### usecase 層

ユースケース層は、ビジネスロジックを呼び出すための層です。具体的には、外部からのリクエストを受け取り、ドメイン層、リポジトリ層を利用してビジネスロジックを実行します。例えば、ユーザーの登録・更新・削除などがあります。

#### delivery 層

デリバリー層は、外部からのリクエストを受け取るための層です。例えば、HTTP、gRPC、REST などのプロトコルを利用して、外部からのリクエストを受け取ります。また、リクエストパラメータのパースやレスポンスの作成などの処理もこの層です。

# application

chat -> schedule
chat -> task management
notification
work_space -> chat -> classification -> summary -> book -> PDF

# technology

Go
Rust
Mysql
Open AI API
localstack

```
- S3
-
```
