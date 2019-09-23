<template>
  <v-container>
    <v-layout text-center wrap>
      <v-flex xs12 mb-5>
        <v-layout justify-center>
          <v-row>
            <v-col>
              <h2>GraphQL Query getting all users called automatically when template is rendered</h2>
              <!-- Apollo watched Graphql query -->
              <ApolloQuery :query="require('../graphql/GetUser.gql')">
                <template slot-scope="{ result: { loading, error, data } }">
                  <!-- Loading -->
                  <div v-if="loading" class="loading apollo">Loading...</div>

                  <!-- Error -->
                  <div v-else-if="error" class="error apollo">An error occured</div>

                  <!-- Result -->
                  <div v-else-if="data" class="result apollo">{{ data }}</div>
                  <!-- No result -->
                  <div v-else class="no-result apollo">No result :(</div>
                </template>
              </ApolloQuery>
              <hr />
              <h2>GraphQL Example Mutation adding new message</h2>
              <ApolloMutation
                :mutation="require('../graphql/AddMessage.gql')"
                :variables="{
                  input: {
                    text: newMessage,
                  },
                }"
                class="form"
                @done="newMessage = ''"
              >
                <template slot-scope="{ mutate }">
                  <form v-on:submit.prevent="formValid && mutate()">
                    <label for="field-message">Message:</label>
                    <input
                      id="field-message"
                      v-model="newMessage"
                      placeholder="Type a message"
                      class="input"
                    />
                  </form>
                </template>
              </ApolloMutation>
              <hr />
              <div class="image-input">
                <label for="field-image">Image</label>
                <input
                  id="field-image"
                  type="file"
                  accept="image/*"
                  required
                  @change="onUploadImage"
                />
              </div>
              <v-sheet height="500">
                <v-calendar :now="today" :value="today" color="primary">
                  <template v-slot:day="{ present, past, date }">
                    <v-row class="fill-height">
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
    name: "Test",
    newMessage: "",
    today: "2019-01-10",
    tracked: {
      "2019-01-09": [23, 45, 10],
      "2019-01-08": [10],
      "2019-01-07": [0, 78, 5],
      "2019-01-06": [0, 0, 50],
      "2019-01-05": [0, 10, 23],
      "2019-01-04": [2, 90],
      "2019-01-03": [10, 32],
      "2019-01-02": [80, 10, 10],
      "2019-01-01": [20, 25, 10]
    },
    colors: ["#1867c0", "#fb8c00", "#000000"],
    category: ["Development", "Meetings", "Slacking"],
    importantLinks: [
      {
        text: "Documentation",
        href: "https://vuetifyjs.com"
      },
      {
        text: "Chat",
        href: "https://community.vuetifyjs.com"
      },
      {
        text: "Made with Vuetify",
        href: "https://madewithvuejs.com/vuetify"
      },
      {
        text: "Twitter",
        href: "https://twitter.com/vuetifyjs"
      },
      {
        text: "Articles",
        href: "https://medium.com/vuetify"
      }
    ],
    whatsNext: [
      {
        text: "Explore components",
        href: "https://vuetifyjs.com/components/api-explorer"
      },
      {
        text: "Select a layout",
        href: "https://vuetifyjs.com/layout/pre-defined"
      },
      {
        text: "Frequently Asked Questions",
        href: "https://vuetifyjs.com/getting-started/frequently-asked-questions"
      }
    ]
  }),
  computed: {
    formValid() {
      return this.newMessage;
    }
  },
  methods: {
    //Different type of example GraphQL Mutation called when onUploadImage method is called
    async onUploadImage({ target }) {
      if (!target.validity.valid) return;
      await this.$apollo.mutate({
        mutation: UPLOAD_FILE,
        variables: {
          file: target.files[0]
        },
        update: (store, { data: { singleUpload } }) => {
          const data = store.readQuery({ query: FILES });
          data.files.push(singleUpload);
          store.writeQuery({ query: FILES, data });
        }
      });
    }
  }
};
</script>
