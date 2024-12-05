<script setup lang="ts">
import { ref, reactive, watch, defineProps } from 'vue';
import VueApexCharts from 'vue3-apexcharts';

type DataPoint = {
    time: number;
    total: number;
    used: number;
    cpuPercent: number;
    net: {
        gbReceived: number;
        gbSent: number;
    };
};


const props = defineProps({
    initialData: {
        type: Array as () => DataPoint[],
        required: true
    }
})

const memChart = reactive({
    series: [
        {
            name: 'Used Memory (GB)',
            data: [] as number[]
        },
    ],
    options: {
        theme: {
            mode: 'dark',
            palette: 'palette1',
            monochrome: {
                enabled: false,
                color: '#255aee',
                shadeTo: 'light',
                shadeIntensity: 0.65
            },
        },
        chart: {
            type: 'line',
            animations: {
                enabled: true,
                easing: 'linear',
                speed: 800,
                dynamicAnimation: {
                    enabled: true,
                    speed: 1000
                }
            },
            toolbar: {
                show: false
            },
            zoom: {
                enabled: false
            }
        },
        title: {
            text: 'Memory Usage',
            align: 'left'
        },
        stroke: {
            curve: 'smooth'
        },
        dataLabels: {
            enabled: false
        },
        noData: {
            text: 'Loading...'
        },
        xaxis: {
            type: 'numeric',
            categories: [] as string[],
            labels: {
                show: true
            },
            tickAmount: 2
        },
        yaxis: {
            tickAmount: 2
        }
    }
});

const cpuChart = reactive({
    series: [
        {
            name: 'CPU Usage (%)',
            data: [] as number[]
        },
    ],
    options: {
        theme: {
            mode: 'dark',
            palette: 'palette1',
            monochrome: {
                enabled: false,
                color: '#255aee',
                shadeTo: 'light',
                shadeIntensity: 0.65
            },
        },
        chart: {
            type: 'line',
            animations: {
                enabled: true,
                easing: 'linear',
                speed: 800,
                dynamicAnimation: {
                    enabled: true,
                    speed: 1000
                }
            },
            toolbar: {
                show: false
            },
            zoom: {
                enabled: false
            }
        },
        title: {
            text: 'CPU Usage',
            align: 'left'
        },
        stroke: {
            curve: 'smooth'
        },
        dataLabels: {
            enabled: false
        },
        noData: {
            text: 'Loading...'
        },
        xaxis: {
            type: 'numeric',
            categories: [] as string[],
            labels: {
                show: true
            },
            tickAmount: 2
        },
        yaxis: {
            tickAmount: 2
        }
    }
});

const networkChart = reactive({
    series: [
        {
            name: 'GB Received',
            data: [] as number[]
        },
        {
            name: 'GB Sent',
            data: [] as number[]
        },
    ],
    options: {
        theme: {
            mode: 'dark',
            palette: 'palette1',
            monochrome: {
                enabled: false,
                color: '#255aee',
                shadeTo: 'light',
                shadeIntensity: 0.65
            },
        },
        chart: {
            type: 'line',
            animations: {
                enabled: true,
                easing: 'linear',
                speed: 800,
                dynamicAnimation: {
                    enabled: true,
                    speed: 1000
                }
            },
            toolbar: {
                show: false
            },
            zoom: {
                enabled: false
            }
        },
        title: {
            text: 'I/O Usage',
            align: 'left'
        },
        stroke: {
            curve: 'smooth'
        },
        dataLabels: {
            enabled: false
        },
        noData: {
            text: 'Loading...'
        },
        xaxis: {
            type: 'numeric',
            categories: [] as string[],
            labels: {
                show: true
            },
            tickAmount: 2
        },
        yaxis: {
            tickAmount: 2
        }
    }
});


watch(() => props.initialData, (newData) => {
    // Обновляем данные графика
    cpuChart.series[0].data = newData.map(item => Math.round(item.cpuPercent));
    memChart.series[0].data = newData.map(item => item.used);
    networkChart.series[0].data = newData.map(item => item.net.gbReceived);
    networkChart.series[1].data = newData.map(item => item.net.gbSent);

    // Обновляем категории по временным меткам
    const newDates = newData.map(item => {
        const date = new Date(item.time * 1000);
        return date.toLocaleString();
    });

    memChart.options.xaxis.categories = newDates
    cpuChart.options.xaxis.categories = newDates
    networkChart.options.xaxis.categories = newDates

}, { immediate: true });
</script>

<template>
    <div class="charts">
        <div class="charts__item">
            <VueApexCharts :options="memChart.options" :series="memChart.series" />
        </div>
        <div class="charts__item">
            <VueApexCharts :options="cpuChart.options" :series="cpuChart.series" />
        </div>
        <div class="charts__item">
            <VueApexCharts :options="networkChart.options" :series="networkChart.series" />
        </div>
    </div>
</template>

<style>
.charts {
    display: flex;
    justify-content: flex-end;
    gap: 20px;
}

.charts:last-child {
    margin-top: 20px;
}

.charts__item {
    width: calc((100%/3));
    height: fit-content;
    color: #000;
}
.vue-apexcharts {
    height: fit-content;
    overflow: hidden;
}
</style>
