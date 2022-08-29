package dataset

// Links to the XLS/XLSX files can be found here: https://www.post.gov.tw/post/internet/Download/all_list.jsp?ID=2201#dl_txt_A09
//go:generate go run sync_xlsx_to_csv.go "https://www.post.gov.tw/post/download/6.5_CEROAD11107.xlsx" "./road.csv"
//go:generate go run sync_xls_to_csv.go "https://www.post.gov.tw/post/download/county_h_10706.xls" "./county.csv"
//go:generate go run sync_xls_to_csv.go "https://www.post.gov.tw/post/download/Village_H_10602.xls" "./village.csv"
