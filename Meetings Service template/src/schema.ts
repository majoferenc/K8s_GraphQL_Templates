import {
  GraphQLObjectType,
  GraphQLInt,
  GraphQLString,
  GraphQLList,
  GraphQLSchema,
  GraphQLNonNull
} from "graphql";

import axios from "axios";
import EntityModel from "../models/entity.model";

const EntityType = new GraphQLObjectType({
  name: "Entity",
  fields: () => ({
    id: { type: GraphQLInt },
    name: { type: GraphQLString }
  })
});

const RootQuery = new GraphQLObjectType({
  name: "RootQueryType",
  fields: {
    entities: {
      type: new GraphQLList(EntityType),
      resolve(parent, args) {
        //TODO implement
        return null;
      }
    },
    entity: {
      type: EntityType,
      args: {
        id: { type: GraphQLInt }
      },
      resolve(parent, args) {
        //TODO implement
        return null;
      }
    }
  }
});

const mutation = new GraphQLObjectType({
  name: "MutationType",
  fields: {
    addEntity: {
      type: EntityType,
      args: {
        id: { type: GraphQLInt },
        name: { type: GraphQLString }
      },
      resolve(parentValue, args) {
        console.log("Adding new Entity");
        return null;
      }
    },
    deleteEntity: {
      type: EntityType,
      args: {
        _id: { type: new GraphQLNonNull(GraphQLString) }
      },
      resolve(parentValue, args) {
        console.log("Deleting Entity");
        return null;
      }
    }
  }
});

export const schema = new GraphQLSchema({
  query: RootQuery,
  mutation
});

module.exports = { schema };
export default { schema };
