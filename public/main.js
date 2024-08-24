document.addEventListener("DOMContentLoaded", () => {
  function setCookie(cname, cvalue, exdays) {
    const d = new Date();
    d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
    let expires = "expires=" + d.toUTCString();
    document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
  }
  async function submitLoginForm(e) {
    e.preventDefault();
    const email = e.target.elements.email.value;
    const password = e.target.elements.password.value;

    const response = await fetch("/api/v1/login", {
      method: "POST",
      body: JSON.stringify({
        email: email,
        password: password,
      }),
    });
    const json = await response.json();

    sessionStorage.setItem("loits_access_token", json.access_token);

    const user = {
      id: json.id,
      email: json.email,
      originalName: json.original_name,
      magicalRating: json.magic_rating,
      birthday: json.birthday,
    };
    sessionStorage.setItem("loits_user", JSON.stringify(user));

    setCookie("loits_access_token", json.access_token, 1);
    return;
  }

  async function submitRegisterForm(e) {
    e.preventDefault();

    const email = e.target.elements.email.value;
    const password = e.target.elements.password.value;
    const firstName = e.target.elements.firstName.value;
    const lastName = e.target.elements.lastName.value;
    const username = e.target.elements.username.value;
    const birthday = new Date().toLocaleString();
    const magicRating = "S";

    if (
      !email || !password || !firstName || !lastName || !username ||
      !birthday || !magicRating
    ) {
      throw new Error("Required field is missing.");
    }

    const response = await fetch("/api/v1/register", {
      method: "POST",
      body: JSON.stringify({
        email,
        password,
        magic_name: username,
        original_name: firstName + " " + lastName,
        magic_rating: magicRating,
        birthday,
      }),
    });

    const json = await response.json();

    const user = {
      id: json.id,
      email: json.email,
      originalName: json.original_name,
      magicalRating: json.magic_rating,
      birthday: json.birthday,
    };
    sessionStorage.setItem("loits_user", JSON.stringify(user));

    setCookie("loits_access_token", json.access_token, 1);
  }

  const loginForm = document.getElementById("loginForm");
  if (loginForm) {
    //On Login Page
    loginForm.addEventListener("submit", submitLoginForm);
  }

  const registerForm = document.getElementById("registerForm");
  if (registerForm) {
    //On Register Page
    registerForm.addEventListener("submit", submitRegisterForm);
  }
});
