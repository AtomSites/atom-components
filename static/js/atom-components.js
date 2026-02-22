(function () {
  "use strict";

  // ============ MODAL ============
  document.addEventListener("click", function (e) {
    // Close button inside modal
    if (e.target.closest("[data-modal-close]")) {
      var modal = e.target.closest(".ac-modal-overlay");
      if (modal) {
        modal.style.display = "none";
      }
      return;
    }
    // Click on overlay background closes modal
    if (e.target.classList.contains("ac-modal-overlay")) {
      e.target.style.display = "none";
    }
  });

  // Escape key closes any open modal
  document.addEventListener("keydown", function (e) {
    if (e.key === "Escape") {
      var modals = document.querySelectorAll('.ac-modal-overlay[style*="flex"]');
      modals.forEach(function (modal) {
        modal.style.display = "none";
      });
    }
  });

  // Global helper to open a modal by ID
  window.acOpenModal = function (id) {
    var modal = document.getElementById(id);
    if (modal) {
      modal.style.display = "flex";
    }
  };

  // Global helper to close a modal by ID
  window.acCloseModal = function (id) {
    var modal = document.getElementById(id);
    if (modal) {
      modal.style.display = "none";
    }
  };

  // ============ TOAST ============
  var toastTimer = 5000;

  window.acToast = function (message, level) {
    level = level || "info";
    var container = document.querySelector(".ac-toast-container");
    if (!container) {
      container = document.createElement("div");
      container.className = "ac-toast-container";
      document.body.appendChild(container);
    }
    var toast = document.createElement("div");
    toast.className = "ac-toast ac-toast-" + level;
    toast.setAttribute("role", "alert");
    toast.innerHTML =
      '<span class="ac-toast-message">' +
      message +
      "</span>" +
      '<button class="ac-toast-close" data-toast-close aria-label="Dismiss">&times;</button>';
    container.appendChild(toast);

    setTimeout(function () {
      toast.classList.add("ac-toast-exit");
      setTimeout(function () {
        toast.remove();
      }, 300);
    }, toastTimer);
  };

  // Toast close button
  document.addEventListener("click", function (e) {
    if (e.target.closest("[data-toast-close]")) {
      var toast = e.target.closest(".ac-toast");
      if (toast) {
        toast.classList.add("ac-toast-exit");
        setTimeout(function () {
          toast.remove();
        }, 300);
      }
    }
  });
})();
