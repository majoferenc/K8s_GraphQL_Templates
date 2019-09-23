/* config file */ 
export const optionsBodyParser = {inflate: true, limit: '1000kb', type: '*/*'};
export const corsOptions = {origin: true, credentials: true};

module.exports = {optionsBodyParser, corsOptions};
export default {optionsBodyParser, corsOptions};