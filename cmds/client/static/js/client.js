var map = L.map('map').setView([35.175,-80.851], 13);

L.tileLayer('https://api.tiles.mapbox.com/v4/{id}/{z}/{x}/{y}.png?access_token=pk.eyJ1IjoibWFwYm94IiwiYSI6ImNpandmbXliNDBjZWd2M2x6bDk3c2ZtOTkifQ._QA7i5Mpkd_m30IGElHziw', {
  maxZoom: 18,
  attribution: 'Map data &copy; <a href="http://openstreetmap.org">OpenStreetMap</a> contributors, ' +
  '<a href="http://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, ' +
  'Imagery © <a href="http://mapbox.com">Mapbox</a>',
  id: 'mapbox.light'
}).addTo(map);



function onEachFeature(feature, layer) {
  // does this feature have a property named note?
  if (feature.properties && feature.properties.note) {
    layer.bindPopup(feature.properties.note);
  }
}

document.addEventListener("DOMContentLoaded", function(){
  // Handler when the DOM is fully loaded
  document.querySelector("#login form").addEventListener("submit", function(ev) {
    ev.preventDefault();

    var email = ev.target[0].value;
    var pwd = ev.target[1].value;

    qwest.post('http://localhost:8080/login', {
      email: email,
      password: pwd
    },
      {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        }
      }
    ).then(function(xhr, response) {
      var token = response.token;
      localStorage.setItem("token", token);
      getItems()
    })

    return false;
  });

  getItems();
});

function getItems() {
  token = localStorage.getItem("token");

  if ( token ) {

    qwest.get(
      'http://localhost:8080/events/014b31ed-ff06-473a-968a-fb44e71846c3/occasions',
      null,
      {
        headers: {
          "Content-Type":"application/json",
          "Authorization":"Bearer " + token
        }
      }
    ).then(function(xhr, response) {

      console.log(response);
      L.geoJSON(response, {onEachFeature: onEachFeature}).addTo(map);
    });
  }
}
