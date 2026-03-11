const reporter = require('cucumber-html-reporter');
const fs = require('fs');

if (!fs.existsSync('reports')) {
    fs.mkdirSync('reports');
}

reporter.generate({
    theme: 'bootstrap',
    jsonFile: 'reports/report.json', 
    output: 'reports/report.html',
    reportSuiteAsScenarios: true,
    launchReport: true,
    metadata: {
        "App Version": "1.0.0",
        "Test Environment": "Development",
    }
});