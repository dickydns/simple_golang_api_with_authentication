require('dotenv').config({ path: '../.env' });

module.exports = {
  default: {
    require: [
      'support/**/*.js',
      'features/step_definitions/**/*.js',
    ],
    paths: ['features/**/**/*.feature'],
    format: [
      'pretty',                           
      'html:reports/report.html',       
      'json:reports/report.json',        
      'junit:reports/report.xml',     
    ]
  }
}