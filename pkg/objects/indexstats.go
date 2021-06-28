package objects

type IndexStats struct {
	AvgDrainRate         float64 `json:"avg_drain_rate"`
	AvgItemSize          float64 `json:"avg_item_size"`
	AvgScanLatency       float64 `json:"avg_scan_latency"`
	CacheHitPercent      float64 `json:"cache_hit_percent"`
	CacheHits            float64 `json:"cache_hits"`
	CacheMisses          float64 `json:"cache_misses"`
	DataSize             float64 `json:"data_size"`
	DiskSize             float64 `json:"disk_size"`
	FragPercent          float64 `json:"frag_percent"`
	InitialBuildProgress float64 `json:"initial_build_progress"`
	ItemsCount           float64 `json:"items_count"`
	LastKnownScanTime    float64 `json:"last_known_scan_time"`
	NumDocsIndexed       float64 `json:"num_docs_indexed"`
	NumDocsPending       float64 `json:"num_docs_pending"`
	NumDocsQueued        float64 `json:"num_docs_queued"`
	NumItemsFlushed      float64 `json:"num_items_flushed"`
	NumPendingRequests   float64 `json:"num_pending_requests"`
	NumRequests          float64 `json:"num_requests"`
	NumRowsReturned      float64 `json:"num_rows_returned"`
	NumScanErrors        float64 `json:"num_scan_errors"`
	NumScanTimeouts      float64 `json:"num_scan_timeouts"`
	RecsInMem            float64 `json:"recs_in_mem"`
	RecsOnDisk           float64 `json:"recs_on_disk"`
	ResidentPercent      float64 `json:"resident_percent"`
	ScanBytesRead        float64 `json:"scan_bytes_read"`
	TotalScanDuration    float64 `json:"total_scan_duration"`
}
