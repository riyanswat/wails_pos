import clipboardy from "clipboardy";
import Swal from "sweetalert2";
import { showAlert, copyToClipboard } from "./utils";
// backend apis:
import {
  Generate,
  Add,
  Delete,
  Search,
  AllData,
  Edit,
} from "../wailsjs/go/backend/App";

class PointOfSale {
  constructor() {
    // input and output elements
    this.passwordElement = document.getElementById("password");
    this.websiteElement = document.getElementById("website");
    this.emailElement = document.getElementById("email");
    this.alertMessage = document.getElementById("alertMessage");
    // buttons
    this.generateBtn = document.getElementById("generate-btn");
    this.addBtn = document.getElementById("add-btn");
    this.searchBtn = document.getElementById("search-btn");
    this.editBtn = document.getElementById("edit-btn");
    this.deleteBtn = document.getElementById("delete-btn");
    this.showAllBtn = document.getElementById("show-btn");
    // values
    this.passwordLength = 8;

    // other elements
    this.dataTableBody = document.getElementById("data-table-body");
    this.allDataElem = document.getElementById("all-data");
    this.appElem = document.getElementById("app");
    this.homeKey = document.getElementById("home");

    this._initEventHandlers();
  }

  // Private methods:

  _initEventHandlers() {
    this.addBtn.addEventListener("click", this.handleAdd.bind(this));
    this.deleteBtn.addEventListener("click", this.handleDelete.bind(this));
    this.generateBtn.addEventListener("click", this.handleGenerate.bind(this));
    this.showAllBtn.addEventListener("click", this.handleShowAll.bind(this));
    this.searchBtn.addEventListener("click", this.handleSearch.bind(this));
    this.editBtn.addEventListener("click", this.handleEdit.bind(this));
    // //? to add shortcuts e.g. ctrl + g for generate:
    // document.addEventListener("keydown", (event) => {
    //   // Check if CTRL + G is pressed
    //   if (event.ctrlKey && event.key === "g") {
    //     // Your action here
    //     this.handleGenerate(this.passwordLength);
    //   }
    // });

    // document.addEventListener("keydown", (event) => {
    //   if (event.key === "Enter") {
    //     Add(
    //       this.websiteElement.value,
    //       this.emailElement.value,
    //       this.passwordElement.value
    //     ).then((res) => {
    //       if (res === "Successful") {
    //         showAlert(this.alertMessage, "Added successfully");
    //         this._clearFields();
    //         return;
    //       } else {
    //         showAlert(this.alertMessage, `${res}`);
    //         return;
    //       }
    //     });
    //   }
    // });
  }

  _clearFields() {
    this.websiteElement.value = "";
    this.emailElement.value = "";
    this.passwordElement.value = "";
  }

  _toggleDisplay() {
    this.allDataElem.style.display = "block";
    this.appElem.style.display = "none";
    this.homeKey.onclick = () => {
      this.allDataElem.style.display = "none";
      this.appElem.style.display = "flex";
    };
  }

  // non-private methods

  async handleAdd() {
    const websiteToAdd = this.websiteElement.value;
    if (!websiteToAdd) {
      showAlert(this.alertMessage, "Please enter a website");
      return;
    }

    try {
      Add(
        this.websiteElement.value,
        this.emailElement.value,
        this.passwordElement.value
      )
        .then((res) => {
          if (res === "Successful") {
            showAlert(this.alertMessage, "Added successfully");
            this._clearFields();
            return;
          } else {
            showAlert(this.alertMessage, `${res}`);
            return;
          }
        })
        .catch((err) => {
          console.error(err);
        });
    } catch (err) {
      console.error(err);
    }
  }

  async handleDelete() {
    const websiteToDelete = this.websiteElement.value;
    if (!websiteToDelete) {
      showAlert(this.alertMessage, "Please enter a website");
      return;
    }

    Search(websiteToDelete).then((res) => {
      if (res[1] == "yes") {
        //* Confirm deletion
        Swal.fire({
          title: "Are you sure?",
          text: `Do you really want to delete '${websiteToDelete}'?`,
          icon: "warning",
          showCancelButton: true,
          confirmButtonColor: "#d33",
          cancelButtonColor: "#3085d6",
          confirmButtonText: "Yes, delete it!",
        }).then((result) => {
          if (result.isConfirmed) {
            //* Deletion:
            Delete(websiteToDelete).then((res) => showAlert(alertMessage, res));

            //* Deletion successful
            Swal.fire({
              title: "Deleted!",
              text: `'${websiteToDelete}' has been deleted.`,
              icon: "success",
            });
            this._clearFields();
          }
        });
      } else if (res[1] == "no") {
        showAlert(this.alertMessage, "Website doesn't exist");
      }
    });
  }

  handleGenerate() {
    Generate(this.passwordLength)
      .then((result) => {
        this.passwordElement.value = result;
        copyToClipboard("Password", result);
      })
      .catch((err) => {
        console.error(err);
      });

    showAlert(alertMessage, "Password generated");
  }

  async handleShowAll() {
    this.dataTableBody.innerHTML = "";

    this._toggleDisplay();
    this._clearFields();

    // this.allDataElem.style.display = "block";
    // this.appElem.style.display = "none";
    // this.homeKey.onclick = () => {
    //   this.allDataElem.style.display = "none";
    //   this.appElem.style.display = "flex";
    // };

    AllData().then((data) => {
      // showAlert(this.alertMessage, data);
      for (let entry of data) {
        // console.log(entry.website);
        // console.log(entry.email);
        // console.log(entry.password);

        const row = document.createElement("tr");
        const websiteCell = document.createElement("td");
        websiteCell.textContent = entry.website;
        const emailCell = document.createElement("td");
        emailCell.textContent = entry.email;
        const passwordCell = document.createElement("td");
        passwordCell.textContent = entry.password;

        row.appendChild(websiteCell);
        row.appendChild(emailCell);
        row.appendChild(passwordCell);
        this.dataTableBody.appendChild(row);
      }
    });
  }

  async handleSearch() {
    const websiteToSearch = this.websiteElement.value;
    if (!websiteToSearch) {
      showAlert(this.alertMessage, "Please enter a website");
      return;
    }

    let itemEmail = "";
    let itemPassword = "";

    Search(websiteToSearch).then((res) => {
      if (res[1] == "no") {
        showAlert(this.alertMessage, "Website doesn't exist");
        return;
      } else {
        itemEmail = res[0].email;
        itemPassword = res[0].password;

        const formattedData = `<strong style="user-select: none;">Email:</strong> ${itemEmail} <span id="copy-email" style="cursor: pointer; user-select: none;">&#x1F4CB;</span>
                    <br><strong style="user-select: none;">Password:</strong> ${itemPassword} <span id="copy-pass" style="cursor: pointer; user-select: none;">&#x1F4CB;</span>`;

        Swal.fire({
          title: websiteToSearch,
          html: formattedData,
          icon: "info",
        });

        let copyEmail = document.getElementById("copy-email");
        let copyPassword = document.getElementById("copy-pass");

        copyEmail.onclick = function () {
          copyToClipboard("email", itemEmail);
          copyEmail.innerHTML = `<span style="background-color: #3498db; color: #fff; padding: 0 10px 0 10px; border-radius: 4px;">Copied!</span>`;
          setTimeout(() => {
            copyEmail.innerHTML = "&#x1F4CB;";
          }, 500);
        };

        copyPassword.onclick = function () {
          copyToClipboard("password", itemPassword);
          copyPassword.innerHTML = `<span style="background-color: #3498db; color: #fff; padding: 0 20px 0 20px; border-radius: 4px;">Copied!</span>`;
          setTimeout(() => {
            copyPassword.innerHTML = "&#x1F4CB;";
          }, 500);
        };

        this._clearFields();
      }
    });
  }

  async handleEdit() {
    if (!this.websiteElement.value) {
      showAlert(this.alertMessage, "Enter a website to edit");
      return;
    }

    let editOption = "";
    const emailHtml =
      '<input id="email-input" class="swal2-input" placeholder="Email">';
    const passwordHtml =
      '<input id="password-input" class="swal2-input" placeholder="Password">';
    const bothHtml = `<input id="email-input" class="swal2-input" placeholder="Enter email"><input id="password-input" class="swal2-input" placeholder="Enter password">`;

    // const editData = {
    //   websiteToEdit: this.websiteElement.value,
    //   newEmail: this.emailElement.value,
    //   newPassword: this.passwordElement.value,
    //   editOption: editOption,
    // };

    const inputOptions = new Promise((resolve) => {
      setTimeout(() => {
        resolve({
          email: "Email",
          password: "Password",
          both: "Both",
        });
      }, 300);
    });

    Swal.fire({
      title: "What do you want to edit?",
      input: "radio",
      inputOptions,
      showCancelButton: true,
      inputValidator: (value) => {
        if (!value) {
          return "You need to choose something!";
        }
      },
    }).then(({ value: option }) => {
      if (option) {
        editOption = option;
        // ? EDIT EMAIL:
        if (editOption == "email") {
          Swal.fire({
            title: "Enter new email",
            html: emailHtml,
            showCancelButton: true,
            confirmButtonText: "Submit",
            cancelButtonText: "Cancel",
            focusConfirm: false,
            preConfirm: () => {
              return document.getElementById("email-input").value;
            },
          }).then((result) => {
            if (result.isConfirmed) {
              console.log(this.websiteElement.value, result.value, "", "email");
              Edit(this.websiteElement.value, result.value, "", "email").then(
                (res) => {
                  showAlert(this.alertMessage, res);
                }
              );
            }
          });
          // ? EDIT PASSWORD
        } else if (editOption == "password") {
          Swal.fire({
            title: "Enter new password",
            html: passwordHtml,
            showCancelButton: true,
            confirmButtonText: "Submit",
            cancelButtonText: "Cancel",
            focusConfirm: false,
            preConfirm: () => {
              return document.getElementById("password-input").value;
            },
          }).then((result) => {
            if (result.isConfirmed) {
              console.log(
                this.websiteElement.value,
                result.value,
                "",
                "password"
              );
              Edit(
                this.websiteElement.value,
                "",
                result.value,
                "password"
              ).then((res) => {
                showAlert(this.alertMessage, res);
              });
              // Swal.fire("You entered:", `${result}`);
            }
          });
          // ? EDIT BOTH
        } else if (editOption == "both") {
          Swal.fire({
            title: "Enter new email and password",
            html: bothHtml,
            showCancelButton: true,
            confirmButtonText: "Submit",
            cancelButtonText: "Cancel",
            focusConfirm: false,
            preConfirm: () => {
              return [
                document.getElementById("email-input").value,
                document.getElementById("password-input").value,
              ];
            },
          }).then((result) => {
            if (result.isConfirmed) {
              console.log(
                this.websiteElement.value,
                result.value[0],
                result.value[1],
                "both"
              );
              Edit(
                this.websiteElement.value,
                result.value[0],
                result.value[1],
                "both"
              ).then((res) => {
                showAlert(this.alertMessage, res);
              });
            }
          });
        }
      }
    });

    // this._clearFields();
  }
}

// instance of PointOfSale
const pos = new PointOfSale();
