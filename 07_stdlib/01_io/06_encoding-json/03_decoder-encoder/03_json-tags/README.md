Decoder 和 Encoder 一样遵循 json 标签规则
，因为底层最终还是调用 `json.Unmarshal` / `json.Marshal` 的逻辑。