// initialize Leaflet
var map = L.map('map').setView({lon: 0, lat: 0}, 2);

// add the OpenStreetMap tiles
L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
  maxZoom: 19,
  attribution: '&copy; <a href="https://openstreetmap.org/copyright">OpenStreetMap contributors</a>'
}).addTo(map);

// show the scale bar on the lower left corner
L.control.scale().addTo(map);

// show a marker on the map
L.marker({lon: 0, lat: 0}).bindPopup('The center of the world').addTo(map);

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
      getEvents();
    })

    return false;
  });

  getEvents();
});

function getEvents() {
  token = localStorage.getItem("token");
  if ( token ) {
    console.log(token);
    qwest.get(
      'http://localhost:8080/events',
      null,
      {
        headers: {
          "Content-Type":"application/json",
          "Authorization":"Bearer " + token
        },
        responseType: 'json'
      }
    ).then(function(xhr, response) {

      var events = response.data;
      try {
        var eventsList = document.getElementById("events");
      } catch(e) {
        console.log(e);
      }
      events.forEach(function(ev) {
        var el = document.createElement("li"),
            a = document.createElement("a");
        a.innerHTML= ev.attributes.name;
        a.href = `/events/${ev.id}`;
        el.appendChild(a);
        eventsList.appendChild(el);
      });

    });
  }
}

function getOccasions() {
  token = localStorage.getItem("token");

  if ( token ) {

    qwest.get(
      'http://localhost:8080/events/014b31ed-ff06-473a-968a-fb44e71846c3/occasions',
      null,
      {
        headers: {
          "Content-Type":"application/json",
          "Authorization":"Bearer " + token
        },
        responseType: 'json'
      }
    ).then(function(xhr, response) {

      console.log(response);
      L.geoJSON(response, {onEachFeature: onEachFeature}).addTo(map);
    });
  }
}
