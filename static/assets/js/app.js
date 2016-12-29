$(document).ready(function(){
    var loc = window.location;
    var uri = 'ws:';

    if (loc.protocol === 'https:') {
      uri = 'wss:';
    }
    uri += '//' + loc.host;
    uri += loc.pathname + 'ws';

    ws = new WebSocket(uri)

    ws.onopen = function() {
      console.log('Connected!');
    }

    ws.onmessage = function(evt) {
      var msg = JSON.parse(event.data);
      console.log(msg)
      switch (msg.type) {
        case "allapps":
            listApps(msg.data)
            break;
        case "appAnalysis":
            console.log(msg.data);
            chart(msg.data);
            break;
      }
    }
    
})

function listApps(apps) {
    for (var i = apps.length - 1; i >= 0; i--) {
        if(i == 0) {
            getAppAnalysis(apps[i].ID)
            elementClasses = "id='"+apps[i].ID+"' class='active'"
        } else
            elementClasses = "id='"+apps[i].ID+"'"
        $( "#applist" ).prepend("<li "+elementClasses+"> <a href='#app-details'> <i class='ti-panel'></i> <p>"+apps[i].Name+"</p> </a> </li>");
    }
}

function getAppAnalysis(appID, duration=-24) {
    var msg = {
        type: "appAnalysis",
        appID: appID,
        duration: duration
        
    };

    ws.send(JSON.stringify(msg));
}

function chart(report){

    labels = []
    series = []

    for (var i = report.RequestsPerMinute.length - 1; i >= 0; i--) {
        for (var key in report.RequestsPerMinute[i]){
            labels.push(key)
            series.push(report.RequestsPerMinute[i][key])
        }
    }
    var data = {
      labels: labels,
      series: [
         series
      ]
    };

    var optionsSales = {
      lineSmooth: false,
      low: 0,
      high: Math.max(series),
      showArea: true,
      height: "245px",
      axisX: {
        showGrid: false,
      },
      lineSmooth: Chartist.Interpolation.simple({
        divisor: 3
      }),
      showLine: true,
      showPoint: false,
    };

    var responsiveSales = [
      ['screen and (max-width: 640px)', {
        axisX: {
          labelInterpolationFnc: function (value) {
            return value[0];
          }
        }
      }]
    ];

    Chartist.Line('#chartRPM', data, optionsSales, responsiveSales);
}