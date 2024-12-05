<template>
  <div class="bg-[#f4f3f2] flex h-screen">
    <!-- Sidebar -->
    <SideBar :current-tab=currentTab @tab-changed="handleTabChange" />

    <!-- Page Content -->
    <div class="flex-1 flex flex-col">
      <!-- Title -->
      <h1
          class="text-transparent bg-gradient-to-r from-purple-400 to-blue-400 bg-clip-text font-bold text-5xl mt-10 ml-20"
      >
        Dashboard
      </h1>

      <!-- First Row -->
      <div class="flex justify-center mt-12 h-[20rem] space-x-96">
        <!-- Total Revenue -->
        <div class="w-[30rem] h-5/6 mt-5 rounded-lg shadow-lg bg-gradient-to-bl from-purple-400/20 to-blue-400/20 p-2 opacity-100">
          <div class="grid grid-rows-1 grid-flow-col gap-4 w-full h-full bg-[#f4f3f2] rounded-lg p-4 overflow-y-auto">
            <div class="col-span-1 w-28 h-28 bg-blue-400/20 rounded-full p-4 mt-[3rem] ml-[2rem] overflow-y-auto text-center text-7xl font-bold">
              $
            </div>
            <div class="col-span-2 row-span-1 text-5xl font-bold mt-[3.5rem]">
              $ {{ userRevenue?.revenue }}
              <br />
              <span class="text-transparent bg-[#5321CA] text-3xl font-semibold  bg-clip-text">Total Revenue</span>
            </div>
          </div>
        </div>
        <!-- Users Onboarded -->
        <div class="w-[30rem] h-5/6 mt-5 rounded-lg shadow-lg bg-gradient-to-br from-purple-400/20 to-blue-400/20 p-2 opacity-100">
          <div class="grid grid-rows-1 grid-flow-col gap-4 w-full h-full bg-[#f4f3f2] rounded-lg p-4 overflow-y-auto">
            <div class="col-span-1 w-28 h-28 bg-purple-400/20 rounded-full p-4 mt-[3rem] ml-[2rem] overflow-y-auto text-center text-7xl font-bold">
              %
            </div>
            <div class="col-span-2 row-span-1 text-5xl font-bold mt-[3.5rem]">
              {{ usersOnboarded?.count }}
              <br />
              <span class="text-transparent bg-[#5321CA] text-2xl font-semibold  bg-clip-text">Vendors Onboarded</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Second Row -->
      <div class="flex justify-center h-[30rem] space-x-28">
        <!-- Activity Metric -->
        <div class="w-[45rem] h-[29rem] mt-5 rounded-lg shadow-lg bg-gradient-to-bl from-purple-400/20 to-blue-400/20 p-2 opacity-100">
          <div class="grid grid-rows-1 grid-flow-col gap-4 w-full h-full bg-[#f4f3f2] rounded-lg p-4 overflow-y-auto">
            <highchart :options="activityOptions" />
          </div>
        </div>
        <!-- Item Sold Metric -->
        <div class="w-[45rem] h-[29rem] mt-5 rounded-lg shadow-lg bg-gradient-to-bl from-purple-400/20 to-blue-400/20 p-2 opacity-100">
          <div class="grid grid-rows-1 grid-flow-col gap-4 w-full h-full bg-[#f4f3f2] rounded-lg p-4 overflow-y-auto">
            <highchart :options="itemsSoldOptions" />
          </div>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRouter } from 'vue-router';
import axios from "axios";
import Cookies from "js-cookie";

interface Revenue {
  revenue: string;
}

interface UsersOnboarded {
  count: number;
}

interface ActivityData {
  period: string;
  number: number;
}

interface Activity {
  data: ActivityData[];
}

interface ItemsData {
  name: string;
  number: number;
}

interface ItemsSold {
  data: ItemsData[];
}

const router = useRouter();
const token = Cookies.get("auth_token")

const userRevenue = ref<Revenue | null>(null);
async function fetchRevenue() {
  try {
    const response = await axios.get("http://185.157.245.42:8080/stats/revenue", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    userRevenue.value = response.data;
  } catch (error) {
    console.error("Error fetching revenue:", error);
  }
}

const usersOnboarded = ref<UsersOnboarded | null>(null);
async function fetchUsersOnboarded() {
  try {
    const response = await axios.get("http://185.157.245.42:8080/stats/users", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    usersOnboarded.value = response.data;
  } catch (error) {
    console.error("Error fetching users:", error);
  }
}

const userActivity = ref<Activity | null>(null);
async function fetchActivity() {
  try {
    const response = await axios.get("http://185.157.245.42:8080/stats/activity", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    userActivity.value = response.data;
  } catch (error) {
    console.error("Error fetching revenue:", error);
  }
}

const userItemsSold = ref<ItemsSold | null>(null);
async function fetchItemsSold() {
  try {
    const response = await axios.get("http://185.157.245.42:8080/stats/items", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    userItemsSold.value = response.data;
  } catch (error) {
    console.error("Error fetching revenue:", error);
  }
}

onMounted(() => {
  fetchRevenue();
  fetchUsersOnboarded();
  fetchActivity();
  fetchItemsSold();
});

const activityOptions = computed(() => {
  if (!userActivity.value || userActivity.value.data.length === 0) {
    return {};
  }
  return {
    chart: {
      type: 'column'
    },
    title: {
      text: 'Vendor Activity'
    },
    accessibility: {
      announceNewData: {
        enabled: true
      }
    },
    xAxis: {
      type: 'category'
    },
    yAxis: {
      title: {
        text: 'Sell number'
      }

    },
    legend: {
      enabled: false
    },
    plotOptions: {
      series: {
        borderWidth: 0,
        dataLabels: {
          enabled: true,
          format: '{point.y:.1f}'
        }
      }
    },

    series: [
      {
        name: 'Sells',
        colorByPoint: false,
        data: [
          {
            name: userActivity.value.data[6].period,
            y: userActivity.value.data[6].number,
          },
          {
            name: userActivity.value.data[5].period,
            y: userActivity.value.data[5].number,
          },
          {
            name: userActivity.value.data[4].period,
            y: userActivity.value.data[4].number,
          },
          {
            name: userActivity.value.data[3].period,
            y: userActivity.value.data[3].number,
          },
          {
            name: userActivity.value.data[2].period,
            y: userActivity.value.data[2].number,
          },
          {
            name: userActivity.value.data[1].period,
            y: userActivity.value.data[1].number,
          },
          {
            name: userActivity.value.data[0].period,
            y: userActivity.value.data[0].number,
          }
        ]
      }
    ],
  };
});

const itemsSoldOptions = computed(() => {
  if (!userItemsSold.value || userItemsSold.value.data.length === 0) {
    return {};
  }
  return {
    chart: {
      type: 'pie',
      custom: {},
      events: {
        render() {
          const chart = this,
              series = chart.series[0];
          let customLabel = chart.options.chart.custom.label;

          if (!customLabel) {
            customLabel = chart.options.chart.custom.label =
                chart.renderer.label(
                    'Total<br/>' +
                    '<strong>2 877 820</strong>'
                )
                    .css({
                      color: '#000',
                      textAnchor: 'middle'
                    })
                    .add();
          }

          const x = series.center[0] + chart.plotLeft,
              y = series.center[1] + chart.plotTop -
                  (customLabel.attr('height') / 2);

          customLabel.attr({
            x,
            y
          });
          // Set font size based on chart diameter
          customLabel.css({
            fontSize: `${series.center[2] / 12}px`
          });
        }
      }
    },
    accessibility: {
      point: {
        valueSuffix: '%'
      }
    },
    title: {
      text: 'Items Sold'
    },
    tooltip: {
      pointFormat: '{series.name}: <b>{point.percentage:.0f}%</b>'
    },
    legend: {
      enabled: false
    },
    plotOptions: {
      series: {
        allowPointSelect: true,
        cursor: 'pointer',
        borderRadius: 8,
        dataLabels: [{
          enabled: true,
          distance: 20,
          format: '{point.name}'
        }, {
          enabled: true,
          distance: -15,
          format: '{point.y:.0f}',
          style: {
            fontSize: '0.9em'
          }
        }],
        showInLegend: true
      }
    },
    series: [{
      name: 'Sells',
      colorByPoint: false,
      innerSize: '75%',

      data: userItemsSold.value.data.map(item => ({
        name: item.name,
        y: item.number,
      })),
    }]
  };
});

const currentTab = ref("Dashboard");

function handleTabChange(newTab: string) {
  currentTab.value = newTab.toLowerCase();
  console.log(`Tab changed to: ${newTab}`);
  router.push(`/merchants/${newTab.toLowerCase()}`);
}
</script>
