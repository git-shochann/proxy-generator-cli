package api

// インスタンス作成 リクエスト時に必要なデータ
type Instance struct {
	Imageid        string // CentOS 7.8
	Instancetypeid string // t2
}
