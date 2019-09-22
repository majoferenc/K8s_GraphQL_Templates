<template>
  <v-container>
    <v-layout
      text-center
      wrap
    >
      <v-flex
        xs12
        mb-5
      >
        <v-layout justify-center>
          <v-row>
            <v-col>
              <v-sheet height="500">
                <v-calendar
                        :now="today"
                        :value="today"
                        color="primary"
                >
                  <template v-slot:day="{ present, past, date }">
                    <v-row
                            class="fill-height"
                    >
                      <template v-if="past && tracked[date]">
                        <v-sheet
                                v-for="(percent, i) in tracked[date]"
                                :key="i"
                                :title="category[i]"
                                :color="colors[i]"
                                :width="`${percent}%`"
                                height="100%"
                                tile
                        ></v-sheet>
                      </template>
                    </v-row>
                  </template>
                </v-calendar>
              </v-sheet>
            </v-col>
          </v-row>
        </v-layout>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  data: () => ({
    today: '2019-01-10',
    tracked: {
      '2019-01-09': [23, 45, 10],
      '2019-01-08': [10],
      '2019-01-07': [0, 78, 5],
      '2019-01-06': [0, 0, 50],
      '2019-01-05': [0, 10, 23],
      '2019-01-04': [2, 90],
      '2019-01-03': [10, 32],
      '2019-01-02': [80, 10, 10],
      '2019-01-01': [20, 25, 10]
    },
    colors: ['#1867c0', '#fb8c00', '#000000'],
    category: ['Development', 'Meetings', 'Slacking'],
    importantLinks: [
      {
        text: 'Documentation',
        href: 'https://vuetifyjs.com'
      },
      {
        text: 'Chat',
        href: 'https://community.vuetifyjs.com'
      },
      {
        text: 'Made with Vuetify',
        href: 'https://madewithvuejs.com/vuetify'
      },
      {
        text: 'Twitter',
        href: 'https://twitter.com/vuetifyjs'
      },
      {
        text: 'Articles',
        href: 'https://medium.com/vuetify'
      }
    ],
    whatsNext: [
      {
        text: 'Explore components',
        href: 'https://vuetifyjs.com/components/api-explorer'
      },
      {
        text: 'Select a layout',
        href: 'https://vuetifyjs.com/layout/pre-defined'
      },
      {
        text: 'Frequently Asked Questions',
        href: 'https://vuetifyjs.com/getting-started/frequently-asked-questions'
      }
    ]
  })
}
</script>
