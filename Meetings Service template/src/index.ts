"use strict";

import * as config from "./express_config";
const app = require("express")();
const graphqlHTTP = require("express-graphql");
const http = require("http").Server(app);
const httpReq = require("http");

const cors = require("cors");
const bodyParser = require("body-parser");

import { schema } from "./schema";

app.use(bodyParser.json());
app.use(bodyParser.raw(config.optionsBodyParser));
app.use(cors(config.corsOptions));

app.use(
  "/graphql",
  graphqlHTTP({
    schema,
    graphiql: true
  })
);

const port = process.env.PORT || 8092;

http.listen(port, () => {
  console.log("Meetings Service launched at ::" + port);
  app.post("/meetings", (request: any, response: any) => {
    //TODO implement
  });
});
