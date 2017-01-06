var map = L.map('map').setView([35.175,-80.851], 13);

L.tileLayer('https://api.tiles.mapbox.com/v4/{id}/{z}/{x}/{y}.png?access_token=pk.eyJ1IjoibWFwYm94IiwiYSI6ImNpandmbXliNDBjZWd2M2x6bDk3c2ZtOTkifQ._QA7i5Mpkd_m30IGElHziw', {
  maxZoom: 18,
  attribution: 'Map data &copy; <a href="http://openstreetmap.org">OpenStreetMap</a> contributors, ' +
  '<a href="http://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, ' +
  'Imagery © <a href="http://mapbox.com">Mapbox</a>',
  id: 'mapbox.light'
}).addTo(map);

qwest.get(
  'http://localhost:8080/events/014b31ed-ff06-473a-968a-fb44e71846c3/occasions',
  null,
  {
    headers: {
      "Content-Type":"application/json",
      "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsImVtYWlsIjoiZ2xlbm4uZ29vZHJpY2hAZ21haWwuY29tIiwiZXhwIjoxNDgzOTc0MjE5LCJpZCI6IjkwNjQyYWUxLTg5NjItNGFiNC1hMjFmLWY5MTQzNzBlMThlNCIsIm5hbWUiOiJHbGVubiBHb29kcmljaCJ9.bz3r8ASO9gC_GRIsevHzWIsPItyESNxkhCAB-Lf9l3I"
    }
  }
).then(function(xhr, response) {

  console.log(response);
  L.geoJSON(response, {onEachFeature: onEachFeature}).addTo(map);
});


function onEachFeature(feature, layer) {
    // does this feature have a property named note?
    if (feature.properties && feature.properties.note) {
        layer.bindPopup(feature.properties.note);
    }
}

