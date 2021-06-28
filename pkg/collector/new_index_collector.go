package collector

import (
	"github.com/arajeet/myexporter/pkg/util"
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

type indexCollector struct {
	m                    MetaCollector
	AvgDrainRate         *prometheus.Desc
	AvgItemSize          *prometheus.Desc
	AvgScanLatency       *prometheus.Desc
	CacheHitPercent      *prometheus.Desc
	CacheHits            *prometheus.Desc
	CacheMisses          *prometheus.Desc
	DataSize             *prometheus.Desc
	DiskSize             *prometheus.Desc
	FragPercent          *prometheus.Desc
	InitialBuildProgress *prometheus.Desc
	ItemsCount           *prometheus.Desc
	LastKnownScanTime    *prometheus.Desc
	NumDocsIndexed       *prometheus.Desc
	NumDocsPending       *prometheus.Desc
	NumDocsQueued        *prometheus.Desc
	NumItemsFlushed      *prometheus.Desc
	NumPendingRequests   *prometheus.Desc
	NumRequests          *prometheus.Desc
	NumRowsReturned      *prometheus.Desc
	NumScanErrors        *prometheus.Desc
	NumScanTimeouts      *prometheus.Desc
	RecsInMem            *prometheus.Desc
	RecsOnDisk           *prometheus.Desc
	ResidentPercent      *prometheus.Desc
	ScanBytesRead        *prometheus.Desc
	TotalScanDuration    *prometheus.Desc
}
type MetaCollector struct {
	mutex  sync.Mutex
	client util.Client

	up             *prometheus.Desc
	scrapeDuration *prometheus.Desc
}

const FQ_NAMESPACE = "couchbase"

func NewIndexCollector(client util.Client) prometheus.Collector {
	const subsystem = "index"
	return &indexCollector{
		m: MetaCollector{
			client: client,
			up: prometheus.NewDesc(
				prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "up"),
				"Couchbase cluster API is responding",
				[]string{"cluster"},
				nil,
			),
			scrapeDuration: prometheus.NewDesc(
				prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "scrape_duration_seconds"),
				"Scrape duration in seconds",
				[]string{"cluster"},
				nil,
			),
		},
		AvgDrainRate: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "avg_drain_rate"),
			"Average Drain Rate of the Index",
			[]string{"Index"},
			nil,
		),
		AvgItemSize: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "avg_item_size"),
			"Average Item Size",
			[]string{"Index"},
			nil,
		),
		AvgScanLatency: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "avg_scan_latency"),
			"Average Scan Latency",
			[]string{"Index"},
			nil,
		),
		CacheHitPercent: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "cache_hit_percent"),
			"Cache Hit Percent",
			[]string{"Index"},
			nil,
		),
		CacheHits: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "cache_hits"),
			"Cache Hits",
			[]string{"Index"},
			nil,
		),
		CacheMisses: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "cache_misses"),
			"Cache Misses",
			[]string{"Index"},
			nil,
		),
		DataSize: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "data_size"),
			"Data Size",
			[]string{"Index"},
			nil,
		),
		DiskSize: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "disk_size"),
			"Disk Size",
			[]string{"Index"},
			nil,
		),
		FragPercent: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "frag_percent"),
			"Frag Percent",
			[]string{"Index"},
			nil,
		),
		InitialBuildProgress: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "initial_build_progress"),
			"Initial Build Progress",
			[]string{"Index"},
			nil,
		),
		ItemsCount: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "items_count"),
			"Items Count",
			[]string{"Index"},
			nil,
		),
		LastKnownScanTime: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "last_known_scan_time"),
			"Last Known Scan Time",
			[]string{"Index"},
			nil,
		),
		NumDocsIndexed: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "num_docs_indexed"),
			"Num Docs Indexed",
			[]string{"Index"},
			nil,
		),
		NumDocsPending: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "num_docs_pending"),
			"Num Docs Pending",
			[]string{"Index"},
			nil,
		),
		NumDocsQueued: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "num_docs_queued"),
			"Num Docs Queued",
			[]string{"Index"},
			nil,
		),
		NumItemsFlushed: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "num_items_flushed"),
			"Num Items Flushed",
			[]string{"Index"},
			nil,
		),
		NumPendingRequests: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "num_pending_requests"),
			"Num Pending Requests",
			[]string{"Index"},
			nil,
		),
		NumRequests: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "num_requests"),
			"Num Requests",
			[]string{"Index"},
			nil,
		),
		NumRowsReturned: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "num_rows_returned"),
			"Num Rows Requests",
			[]string{"Index"},
			nil,
		),
		NumScanErrors: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "num_scan_errors"),
			"Num Scan Errors",
			[]string{"Index"},
			nil,
		),
		NumScanTimeouts: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "num_scan_timeouts"),
			"Num Scan Timeouts",
			[]string{"Index"},
			nil,
		),
		RecsInMem: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "recs_in_mem"),
			"Recs in mem",
			[]string{"Index"},
			nil,
		),
		RecsOnDisk: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "recs_on_disk"),
			"Recs on Disk",
			[]string{"Index"},
			nil,
		),
		ResidentPercent: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "resident_percent"),
			"Resident Percent",
			[]string{"Index"},
			nil,
		),
		ScanBytesRead: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "scan_bytes_read"),
			"Scan Bytes Read",
			[]string{"Index"},
			nil,
		),
		TotalScanDuration: prometheus.NewDesc(
			prometheus.BuildFQName(FQ_NAMESPACE+subsystem, "", "total_scan_duration"),
			"Total Scan Duration",
			[]string{"Index"},
			nil,
		),
	}
}
func (i *indexCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- i.AvgDrainRate
	ch <- i.AvgItemSize
	ch <- i.AvgScanLatency
	ch <- i.CacheHitPercent
	ch <- i.CacheHits
	ch <- i.CacheMisses
	ch <- i.DataSize
	ch <- i.DiskSize
	ch <- i.FragPercent
	ch <- i.InitialBuildProgress
	ch <- i.ItemsCount
	ch <- i.LastKnownScanTime
	ch <- i.NumDocsIndexed
	ch <- i.NumDocsPending
	ch <- i.NumDocsQueued
	ch <- i.NumItemsFlushed
	ch <- i.NumPendingRequests
	ch <- i.NumRequests
	ch <- i.NumRowsReturned
	ch <- i.NumScanErrors
	ch <- i.NumScanTimeouts
	ch <- i.RecsInMem
	ch <- i.RecsOnDisk
	ch <- i.ResidentPercent
	ch <- i.ScanBytesRead
	ch <- i.TotalScanDuration
}

func (i *indexCollector) Collect(ch chan<- prometheus.Metric) {
	{

		m := util.CallIndexstats(i.m.client)
		for k, v := range m {
			//bucketnameIndex := strings.Split(k, ":")
			ch <- prometheus.MustNewConstMetric(i.AvgDrainRate, prometheus.GaugeValue, v.AvgDrainRate, k)
			ch <- prometheus.MustNewConstMetric(i.AvgItemSize, prometheus.GaugeValue, v.AvgItemSize, k)
			ch <- prometheus.MustNewConstMetric(i.AvgScanLatency, prometheus.GaugeValue, v.AvgScanLatency, k)
			ch <- prometheus.MustNewConstMetric(i.CacheHitPercent, prometheus.GaugeValue, v.CacheHitPercent, k)
			ch <- prometheus.MustNewConstMetric(i.CacheHits, prometheus.GaugeValue, v.CacheHits, k)
			ch <- prometheus.MustNewConstMetric(i.CacheMisses, prometheus.GaugeValue, v.CacheMisses, k)
			ch <- prometheus.MustNewConstMetric(i.DataSize, prometheus.GaugeValue, v.DataSize, k)
			ch <- prometheus.MustNewConstMetric(i.DiskSize, prometheus.GaugeValue, v.DiskSize, k)
			ch <- prometheus.MustNewConstMetric(i.FragPercent, prometheus.GaugeValue, v.FragPercent, k)
			ch <- prometheus.MustNewConstMetric(i.InitialBuildProgress, prometheus.GaugeValue, v.InitialBuildProgress, k)
			ch <- prometheus.MustNewConstMetric(i.ItemsCount, prometheus.GaugeValue, v.ItemsCount, k)
			ch <- prometheus.MustNewConstMetric(i.LastKnownScanTime, prometheus.GaugeValue, v.LastKnownScanTime, k)
			ch <- prometheus.MustNewConstMetric(i.NumDocsIndexed, prometheus.GaugeValue, v.NumDocsIndexed, k)
			ch <- prometheus.MustNewConstMetric(i.NumDocsPending, prometheus.GaugeValue, v.NumDocsPending, k)
			ch <- prometheus.MustNewConstMetric(i.NumDocsQueued, prometheus.GaugeValue, v.NumDocsQueued, k)
			ch <- prometheus.MustNewConstMetric(i.NumItemsFlushed, prometheus.GaugeValue, v.NumItemsFlushed, k)
			ch <- prometheus.MustNewConstMetric(i.NumPendingRequests, prometheus.GaugeValue, v.NumPendingRequests, k)
			ch <- prometheus.MustNewConstMetric(i.NumRequests, prometheus.GaugeValue, v.NumRequests, k)
			ch <- prometheus.MustNewConstMetric(i.NumRowsReturned, prometheus.GaugeValue, v.NumRowsReturned, k)
			ch <- prometheus.MustNewConstMetric(i.NumScanErrors, prometheus.GaugeValue, v.NumScanErrors, k)
			ch <- prometheus.MustNewConstMetric(i.NumScanTimeouts, prometheus.GaugeValue, v.NumScanTimeouts, k)
			ch <- prometheus.MustNewConstMetric(i.RecsInMem, prometheus.GaugeValue, v.RecsInMem, k)
			ch <- prometheus.MustNewConstMetric(i.RecsOnDisk, prometheus.GaugeValue, v.RecsOnDisk, k)
			ch <- prometheus.MustNewConstMetric(i.ResidentPercent, prometheus.GaugeValue, v.ResidentPercent, k)
			ch <- prometheus.MustNewConstMetric(i.ScanBytesRead, prometheus.GaugeValue, v.ScanBytesRead, k)
			ch <- prometheus.MustNewConstMetric(i.TotalScanDuration, prometheus.GaugeValue, v.TotalScanDuration, k)
		}

	}
}
