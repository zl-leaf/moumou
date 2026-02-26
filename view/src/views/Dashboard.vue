<template>
  <div class="dashboard">
    <div class="dashboard-header">
      <div>
        <h1 class="dashboard-title">欢迎使用 Moumou Admin</h1>
        <p class="dashboard-desc">
          这里是系统仪表盘，你可以在左侧菜单进入账号、角色、权限和内容等管理功能。
        </p>
      </div>
    </div>

    <div class="dashboard-grid">
      <div class="card">
        <div class="card-header">
          <span class="card-title">近 7 日访问量</span>
        </div>
        <v-chart class="chart" :option="barOption" autoresize />
      </div>

      <div class="card">
        <div class="card-header">
          <span class="card-title">近 7 日活跃用户趋势</span>
        </div>
        <v-chart class="chart" :option="lineOption" autoresize />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { use } from 'echarts/core'
import { BarChart, LineChart } from 'echarts/charts'
import { GridComponent, TooltipComponent, LegendComponent } from 'echarts/components'
import { CanvasRenderer } from 'echarts/renderers'
import VChart from 'vue-echarts'

use([BarChart, LineChart, GridComponent, TooltipComponent, LegendComponent, CanvasRenderer])

const barOption = ref({
  tooltip: {
    trigger: 'axis',
  },
  grid: {
    left: 40,
    right: 20,
    top: 40,
    bottom: 40,
  },
  xAxis: {
    type: 'category',
    data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日'],
  },
  yAxis: {
    type: 'value',
  },
  series: [
    {
      name: '访问量',
      type: 'bar',
      data: [120, 200, 150, 80, 70, 110, 130],
      itemStyle: {
        color: '#1677ff',
      },
      barWidth: 24,
    },
  ],
})

const lineOption = ref({
  tooltip: {
    trigger: 'axis',
  },
  legend: {
    data: ['活跃用户', '新注册'],
  },
  grid: {
    left: 40,
    right: 20,
    top: 40,
    bottom: 40,
  },
  xAxis: {
    type: 'category',
    data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日'],
  },
  yAxis: {
    type: 'value',
  },
  series: [
    {
      name: '活跃用户',
      type: 'line',
      smooth: true,
      data: [30, 52, 45, 60, 75, 68, 80],
      itemStyle: {
        color: '#52c41a',
      },
      areaStyle: {
        color: 'rgba(82,196,26,0.15)',
      },
    },
    {
      name: '新注册',
      type: 'line',
      smooth: true,
      data: [10, 18, 12, 20, 24, 22, 26],
      itemStyle: {
        color: '#faad14',
      },
    },
  ],
})
</script>

<style scoped>
.dashboard {
  padding: 16px 20px;
}

.dashboard-header {
  margin-bottom: 16px;
}

.dashboard-title {
  font-size: 22px;
  font-weight: 600;
  margin-bottom: 4px;
}

.dashboard-desc {
  color: #8c8c8c;
  font-size: 13px;
}

.dashboard-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 16px;
}

.card {
  background: #fff;
  border-radius: 12px;
  padding: 16px 16px 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.03);
}

.card-header {
  margin-bottom: 8px;
}

.card-title {
  font-size: 14px;
  font-weight: 500;
}

.chart {
  width: 100%;
  height: 260px;
}

@media (max-width: 1024px) {
  .dashboard-grid {
    grid-template-columns: 1fr;
  }
}
</style>
