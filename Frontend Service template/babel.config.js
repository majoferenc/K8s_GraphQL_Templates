module.exports = {
  presets: [
    '@vue/app'
  ],
  plugins: [
    ['transform-imports', {
      'vuetify': {
        // eslint-disable-next-line no-undef
        'transform': `vuetify/es5/components/${member}`,
        'preventFullImport': true
      }
    }]
  ]
}
