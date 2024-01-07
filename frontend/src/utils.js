import clipboardy from "clipboardy";

// import "./style.css";
// import "./app.css";

// const alertMessage = document.getElementById("alertMessage");

export function showAlert(elem, text) {
  // elem = document.getElementById("alertMessage");
  elem.style.display = "inline-block";
  elem.innerText = text;

  setTimeout(() => {
    elem.style.display = "none";
  }, 1400);
}

export function copyToClipboard(el, val) {
  clipboardy.write(val, function (err) {
    if (err) {
      console.error(err);
    } else {
      console.log(`${el} copied to clipboard!`);
    }
  });
}

// end
// if (typeof module !== "undefined") {
//   module.exports = showAlert;
// }
