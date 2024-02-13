let currentForm = "form1";
function toggleForms() {
const form1 = document.getElementById("form1");
const form2 = document.getElementById("form2");

if (currentForm === "form1") {
    form1.style.display = "none";
    form2.style.display = "block";
    currentForm = "form2";
} else {
    form1.style.display = "block";
    form2.style.display = "none"; 
    currentForm = "form1";
        }
    }
