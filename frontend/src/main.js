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
  PasswordProtect,
} from "../wailsjs/go/backend/App";

class PointOfSale {
  constructor() {
    // main app password from the backend:
    // In 'backend/utils.go', I have:
    // var AppPassword string = "riyan"
    // which is the password of the main app

    const storedPassword = sessionStorage.getItem("appPassword");
    if (storedPassword) {
      this.passwordProtection = storedPassword;
    } else {
      // If no stored password, fetch from the backend
      PasswordProtect().then((main_password) => {
        this.passwordProtection = main_password; // set it to 'main_password' to enable pass protection
      });
    }

    // ----------------------------------------

    // input and output elements
    this.priceElement = document.getElementById("price");
    this.itemElement = document.getElementById("item");
    this.quantityElement = document.getElementById("quantity");
    this.alertMessage = document.getElementById("alertMessage");
    // buttons
    this.generateBtn = document.getElementById("generate-btn");
    this.addBtn = document.getElementById("add-btn");
    this.searchBtn = document.getElementById("search-btn");
    this.editBtn = document.getElementById("edit-btn");
    this.deleteBtn = document.getElementById("delete-btn");
    this.showAllBtn = document.getElementById("show-btn");
    // values
    this.priceLength = 8;

    // other elements
    this.dataTableBody = document.getElementById("data-table-body");
    this.allDataElem = document.getElementById("all-data");
    this.appElem = document.getElementById("app");
    this.homeKey = document.getElementById("home");

    this._initEventHandlers();
  }

  // ------------------------------------------------------------
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
    //     this.handleGenerate(this.priceLength);
    //   }
    // });

    // document.addEventListener("keydown", (event) => {
    //   if (event.key === "Enter") {
    //     Add(
    //       this.itemElement.value,
    //       this.quantityElement.value,
    //       this.priceElement.value
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

  // ------------------------------------------------------------

  _clearFields() {
    this.itemElement.value = "";
    this.quantityElement.value = "";
    this.priceElement.value = "";
  }

  // ------------------------------------------------------------

  _toggleDisplay() {
    this.allDataElem.style.display = "block";
    this.appElem.style.display = "none";
    this.homeKey.onclick = () => {
      this.allDataElem.style.display = "none";
      this.appElem.style.display = "flex";
    };
  }

  // ------------------------------------------------------------

  async _getPassword() {
    // Check if the password is already stored in sessionStorage
    const storedPassword = sessionStorage.getItem("appPassword");

    if (storedPassword) {
      return storedPassword;
    }

    // If not stored, prompt the user
    const { value: password } = await Swal.fire({
      title: "Enter your password",
      input: "password",
      inputLabel: "Password",
      inputPlaceholder: "Enter your password",
      inputAttributes: {
        maxlength: "10",
        autocapitalize: "off",
        autocorrect: "off",
      },
    });

    // Store the entered password in sessionStorage
    sessionStorage.setItem("appPassword", password);

    return password;
  }

  // ------------------------------------------------------------

  async _validatePassword() {
    // Check if password protection is enabled
    if (this.passwordProtection !== null) {
      const storedPassword = sessionStorage.getItem("appPassword");

      if (!storedPassword) {
        // If no stored password, validate the entered password
        const enteredPassword = await this._getPassword();

        if (enteredPassword !== this.passwordProtection) {
          Swal.fire("Incorrect password");
          return false;
        }
      } else {
        // If a password is stored, validate it against the backend password
        if (storedPassword !== this.passwordProtection) {
          Swal.fire("Incorrect password");
          return false;
        }
      }
    }
    return true;
  }

  // ------------------------------------------------------------

  // non-private methods

  async handleAdd() {
    const itemToAdd = this.itemElement.value;
    const quantityToAdd = this.quantityElement.value;
    const priceToAdd = this.priceElement.value;

    if (!itemToAdd || !quantityToAdd || !priceToAdd) {
      showAlert(this.alertMessage, "Please fill all the fields");
      return;
    }

    // validate password
    if (!(await this._validatePassword())) {
      return;
    }

    Add(
      this.itemElement.value,
      this.quantityElement.value,
      this.priceElement.value
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
  }

  // ------------------------------------------------------------

  async handleDelete() {
    const itemToDelete = this.itemElement.value;
    if (!itemToDelete) {
      showAlert(this.alertMessage, "Please enter an item");
      return;
    }

    // validate password
    if (!(await this._validatePassword())) {
      return;
    }

    Search(itemToDelete).then((res) => {
      if (res[1] == "yes") {
        //* Confirm deletion

        Swal.fire({
          title: "Are you sure?",
          text: `Do you really want to delete '${itemToDelete}'?`,
          icon: "warning",
          showCancelButton: true,
          confirmButtonColor: "#d33",
          cancelButtonColor: "#3085d6",
          confirmButtonText: "Yes, delete it!",
        }).then((result) => {
          if (result.isConfirmed) {
            //* Deletion:
            Delete(itemToDelete).then((res) => showAlert(alertMessage, res));

            //* Deletion successful
            Swal.fire({
              title: "Deleted!",
              text: `'${itemToDelete}' has been deleted.`,
              icon: "success",
            });
            this._clearFields();
          }
        });
      } else if (res[1] == "no") {
        showAlert(this.alertMessage, "Item doesn't exist");
      }
    });
  }

  // ------------------------------------------------------------

  handleGenerate() {
    Generate(this.priceLength)
      .then((result) => {
        this.priceElement.value = result;
        copyToClipboard("Price", result);
      })
      .catch((err) => {
        console.error(err);
      });

    showAlert(alertMessage, "Price generated");
  }

  // ------------------------------------------------------------

  async handleShowAll() {
    // validate password
    if (!(await this._validatePassword())) {
      return;
    }
    this.dataTableBody.innerHTML = "";

    this._toggleDisplay();
    this._clearFields();

    AllData().then((data) => {
      console.log(data);
      // showAlert(this.alertMessage, data);
      for (let entry of data) {
        const row = document.createElement("tr");
        const itemCell = document.createElement("td");
        itemCell.textContent = entry.item;
        const quantityCell = document.createElement("td");
        quantityCell.textContent = entry.quantity;
        const priceCell = document.createElement("td");
        priceCell.textContent = entry.price;

        row.appendChild(itemCell);
        row.appendChild(quantityCell);
        row.appendChild(priceCell);
        this.dataTableBody.appendChild(row);
      }
    });
  }

  // ------------------------------------------------------------

  async handleSearch() {
    const itemToSearch = this.itemElement.value;
    if (!itemToSearch) {
      showAlert(this.alertMessage, "Please enter an item");
      return;
    }

    // validate password
    if (!(await this._validatePassword())) {
      return;
    }

    let itemQuantity = "";
    let itemPrice = "";

    Search(itemToSearch).then((res) => {
      if (res[1] == "no") {
        showAlert(this.alertMessage, "Item doesn't exist");
        return;
      } else {
        itemQuantity = res[0].quantity;
        itemPrice = res[0].price;

        const formattedData = `<strong style="user-select: none;">Quantity:</strong> ${itemQuantity} <span id="copy-quantity" style="cursor: pointer; user-select: none;">&#x1F4CB;</span>
                    <br><strong style="user-select: none;">Price:</strong> ${itemPrice} <span id="copy-pass" style="cursor: pointer; user-select: none;">&#x1F4CB;</span>`;

        Swal.fire({
          title: itemToSearch,
          html: formattedData,
          icon: "info",
        });

        let copyQuantity = document.getElementById("copy-quantity");
        let copyPrice = document.getElementById("copy-pass");

        copyQuantity.onclick = function () {
          copyToClipboard("quantity", itemQuantity);
          copyQuantity.innerHTML = `<span style="background-color: #3498db; color: #fff; padding: 0 10px 0 10px; border-radius: 4px;">Copied!</span>`;
          setTimeout(() => {
            copyQuantity.innerHTML = "&#x1F4CB;";
          }, 500);
        };

        copyPrice.onclick = function () {
          copyToClipboard("price", itemPrice);
          copyPrice.innerHTML = `<span style="background-color: #3498db; color: #fff; padding: 0 20px 0 20px; border-radius: 4px;">Copied!</span>`;
          setTimeout(() => {
            copyPrice.innerHTML = "&#x1F4CB;";
          }, 500);
        };

        this._clearFields();
      }
    });
  }

  // ------------------------------------------------------------

  async handleEdit() {
    let itemToEdit = this.itemElement.value;
    let spacesRegex = /^\s+$/;
    if (!itemToEdit || spacesRegex.test(itemToEdit)) {
      showAlert(this.alertMessage, "Enter an item to edit");
      return;
    }

    // validate password
    if (!(await this._validatePassword())) {
      return;
    }

    Search(itemToEdit).then((res) => {
      if (res[1] == "no") {
        // ------------------ ITEM DOESN'T EXIST: ------------------
        console.log("Item doesn't exist");
        showAlert(this.alertMessage, "Item doesn't exist");
        return;
      } else {
        // ------------------ ITEM EXISTS: ------------------
        console.log("Item exists");

        let editOption = "";
        const quantityHtml =
          '<input id="quantity-input" class="swal2-input" placeholder="Quantity">';
        const priceHtml =
          '<input id="price-input" class="swal2-input" placeholder="Price">';
        const bothHtml = `<input id="quantity-input" class="swal2-input" placeholder="Enter quantity"><input id="price-input" class="swal2-input" placeholder="Enter price">`;

        const inputOptions = new Promise((resolve) => {
          setTimeout(() => {
            resolve({
              quantity: "Quantity",
              price: "Price",
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
            // ? -------------------- EDIT QUANTITY --------------------

            if (editOption == "quantity") {
              Swal.fire({
                title: "Enter new quantity",
                html: quantityHtml,
                showCancelButton: true,
                confirmButtonText: "Submit",
                cancelButtonText: "Cancel",
                focusConfirm: false,
                preConfirm: () => {
                  return document.getElementById("quantity-input").value;
                },
              }).then((result) => {
                if (result.isConfirmed) {
                  console.log(
                    this.itemElement.value,
                    result.value,
                    "",
                    "quantity"
                  );
                  Edit(
                    this.itemElement.value,
                    result.value,
                    "",
                    "quantity"
                  ).then((res) => {
                    showAlert(this.alertMessage, res);
                  });
                }
              });

              // ? -------------------- EDIT PRICE --------------------
            } else if (editOption == "price") {
              Swal.fire({
                title: "Enter new price",
                html: priceHtml,
                showCancelButton: true,
                confirmButtonText: "Submit",
                cancelButtonText: "Cancel",
                focusConfirm: false,
                preConfirm: () => {
                  return document.getElementById("price-input").value;
                },
              }).then((result) => {
                if (result.isConfirmed) {
                  console.log(
                    this.itemElement.value,
                    result.value,
                    "",
                    "price"
                  );
                  Edit(this.itemElement.value, "", result.value, "price").then(
                    (res) => {
                      showAlert(this.alertMessage, res);
                    }
                  );
                }
              });
              // ? -------------------- EDIT BOTH --------------------
            } else if (editOption == "both") {
              Swal.fire({
                title: "Enter new quantity and price",
                html: bothHtml,
                showCancelButton: true,
                confirmButtonText: "Submit",
                cancelButtonText: "Cancel",
                focusConfirm: false,
                preConfirm: () => {
                  return [
                    document.getElementById("quantity-input").value,
                    document.getElementById("price-input").value,
                  ];
                },
              }).then((result) => {
                if (result.isConfirmed) {
                  console.log(
                    this.itemElement.value,
                    result.value[0],
                    result.value[1],
                    "both"
                  );
                  Edit(
                    this.itemElement.value,
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
      }
    });

    // this._clearFields();
  }

  // ------------------------------------------------------------
}

// instance of PointOfSale
const pos = new PointOfSale();
