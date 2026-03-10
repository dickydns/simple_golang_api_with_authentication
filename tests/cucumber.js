require('dotenv').config();

module.exports = {
  default: {
    require: [
      'support/**/*.js',
      'features/step_definitions/**/*.js',
    ],
    paths: ['features/**/**/*.feature'],
    format: ['pretty'] 
  }
}