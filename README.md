# Golang 後端工程師筆試題目

**！嚴格禁止發送合併請求至本倉庫！**

## 注意事項

- 本題不允許使用第三方套件
- **三道題目都已在各自的資料夾中**
- 三道題目請在不同的 git branch 中完成
- 全部完成後請將三個 git branch 合併至 main

## 題目一

[題目一](assembly_line/main.go)

請模擬流水線, 五個員工處理三種物品

1. 三種物品數量個十件
2. 三種物品處理時間需不一樣
3. 物品的處理順序請隨機打亂
4. 物品處理需透過 interface 來傳遞
5. 每個員工一次只能處理一種物品
6. 開始以及結束處理都需要打印紀錄
7. 統計總處理時間, 及每個員工處理了多少物品

## 題目二

[題目二](swap/main.go)

請完成 swap 函式, 交換兩個變數的值

1. 允許panic但必須是顯式調用
2. 地址不允許改變
3. 禁止修改 swap 類型標記

### 範例

![swap.png](images/swap.png)

## 題目三

[題目三](trim_all_strings/main.go)

請完成 TrimAllStrings 函式, 移除所有空白字元

1. 禁止修改 TrimAllStrings 類型標記

### 範例

![trim_all_strings.png](images/trim_all_strings.png)
