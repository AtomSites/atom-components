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
  // ============ DATEPICKER ============
  var MONTHS = [
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December",
  ];
  var MONTHS_SHORT = [
    "Jan",
    "Feb",
    "Mar",
    "Apr",
    "May",
    "Jun",
    "Jul",
    "Aug",
    "Sep",
    "Oct",
    "Nov",
    "Dec",
  ];
  var WEEKDAYS = ["Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"];

  var datepickers = {};

  function dpGetState(root) {
    var id = root.id;
    if (!datepickers[id]) {
      var now = new Date();
      var minYear = parseInt(root.dataset.acDatepickerMinYear, 10) || now.getFullYear() - 100;
      var maxYear = parseInt(root.dataset.acDatepickerMaxYear, 10) || now.getFullYear() + 20;
      var hiddenInput = root.querySelector("[data-ac-datepicker-value]");
      var existing = hiddenInput ? hiddenInput.value : "";
      var selYear = null;
      var selMonth = null;
      var selDay = null;
      if (existing) {
        var parts = existing.split("-");
        selYear = parseInt(parts[0], 10);
        selMonth = parseInt(parts[1], 10) - 1;
        selDay = parseInt(parts[2], 10);
      }
      datepickers[id] = {
        step: "year",
        minYear: minYear,
        maxYear: maxYear,
        viewYear: selYear || now.getFullYear(),
        viewMonth: selMonth !== null ? selMonth : now.getMonth(),
        selYear: selYear,
        selMonth: selMonth,
        selDay: selDay,
      };
    }
    return datepickers[id];
  }

  function dpRenderYearGrid(state) {
    var html = '<div class="ac-datepicker-year-grid">';
    var now = new Date();
    for (var y = state.minYear; y <= state.maxYear; y++) {
      var cls = "ac-datepicker-cell";
      if (y === state.selYear) cls += " ac-datepicker-cell-selected";
      if (y === now.getFullYear()) cls += " ac-datepicker-cell-today";
      html += '<div class="' + cls + '" data-ac-datepicker-year="' + y + '">' + y + "</div>";
    }
    html += "</div>";
    return html;
  }

  function dpRenderMonthGrid(state) {
    var html = '<div class="ac-datepicker-month-grid">';
    var now = new Date();
    for (var m = 0; m < 12; m++) {
      var cls = "ac-datepicker-cell";
      if (state.selYear === state.viewYear && m === state.selMonth)
        cls += " ac-datepicker-cell-selected";
      if (state.viewYear === now.getFullYear() && m === now.getMonth())
        cls += " ac-datepicker-cell-today";
      html += '<div class="' + cls + '" data-ac-datepicker-month="' + m + '">' + MONTHS_SHORT[m] + "</div>";
    }
    html += "</div>";
    return html;
  }

  function dpRenderDayGrid(state) {
    var html = '<div class="ac-datepicker-day-grid">';
    var now = new Date();
    var todayStr = now.getFullYear() + "-" + (now.getMonth() + 1) + "-" + now.getDate();

    for (var d = 0; d < 7; d++) {
      html += '<div class="ac-datepicker-weekday">' + WEEKDAYS[d] + "</div>";
    }

    var firstDay = new Date(state.viewYear, state.viewMonth, 1).getDay();
    var daysInMonth = new Date(state.viewYear, state.viewMonth + 1, 0).getDate();
    var prevDays = new Date(state.viewYear, state.viewMonth, 0).getDate();

    // Previous month fill
    for (var p = firstDay - 1; p >= 0; p--) {
      var pd = prevDays - p;
      html +=
        '<div class="ac-datepicker-cell ac-datepicker-cell-other" data-ac-datepicker-day-other="prev-' +
        pd +
        '">' +
        pd +
        "</div>";
    }

    // Current month days
    for (var i = 1; i <= daysInMonth; i++) {
      var cls = "ac-datepicker-cell";
      var dayStr = state.viewYear + "-" + (state.viewMonth + 1) + "-" + i;
      if (state.selYear === state.viewYear && state.selMonth === state.viewMonth && state.selDay === i)
        cls += " ac-datepicker-cell-selected";
      if (dayStr === todayStr) cls += " ac-datepicker-cell-today";
      html += '<div class="' + cls + '" data-ac-datepicker-day="' + i + '">' + i + "</div>";
    }

    // Next month fill (42 cells total = 6 rows)
    var totalCells = firstDay + daysInMonth;
    var remaining = 42 - totalCells;
    for (var n = 1; n <= remaining; n++) {
      html +=
        '<div class="ac-datepicker-cell ac-datepicker-cell-other" data-ac-datepicker-day-other="next-' +
        n +
        '">' +
        n +
        "</div>";
    }

    html += "</div>";
    return html;
  }

  function dpRender(root) {
    var state = dpGetState(root);
    var body = root.querySelector("[data-ac-datepicker-body]");
    var title = root.querySelector("[data-ac-datepicker-title]");
    var backBtn = root.querySelector("[data-ac-datepicker-back]");

    if (state.step === "year") {
      title.textContent = "Select Year";
      backBtn.style.visibility = "hidden";
      body.innerHTML = dpRenderYearGrid(state);
      // Auto-scroll to selected or current year
      var yearGrid = body.querySelector(".ac-datepicker-year-grid");
      var selected =
        yearGrid.querySelector(".ac-datepicker-cell-selected") ||
        yearGrid.querySelector(".ac-datepicker-cell-today");
      if (selected) {
        selected.scrollIntoView({ block: "center", behavior: "instant" });
      }
    } else if (state.step === "month") {
      title.textContent = state.viewYear.toString();
      backBtn.style.visibility = "visible";
      body.innerHTML = dpRenderMonthGrid(state);
    } else if (state.step === "day") {
      title.textContent = MONTHS[state.viewMonth] + " " + state.viewYear;
      backBtn.style.visibility = "visible";
      body.innerHTML = dpRenderDayGrid(state);
    }
  }

  function dpOpen(root) {
    var state = dpGetState(root);
    var overlay = root.querySelector(".ac-datepicker-overlay");
    var trigger = root.querySelector("[data-ac-datepicker-trigger]");
    if (state.selYear !== null) {
      state.step = "day";
      state.viewYear = state.selYear;
      state.viewMonth = state.selMonth;
    } else {
      state.step = "year";
    }
    dpRender(root);
    overlay.style.display = "flex";
    trigger.setAttribute("aria-expanded", "true");
  }

  function dpClose(root) {
    var overlay = root.querySelector(".ac-datepicker-overlay");
    var trigger = root.querySelector("[data-ac-datepicker-trigger]");
    overlay.style.display = "none";
    trigger.setAttribute("aria-expanded", "false");
  }

  function dpConfirm(root) {
    var state = dpGetState(root);
    if (state.selYear === null || state.selMonth === null || state.selDay === null) return;
    var mm = String(state.selMonth + 1).padStart(2, "0");
    var dd = String(state.selDay).padStart(2, "0");
    var iso = state.selYear + "-" + mm + "-" + dd;
    var display = MONTHS_SHORT[state.selMonth] + " " + state.selDay + ", " + state.selYear;

    var hidden = root.querySelector("[data-ac-datepicker-value]");
    var trigger = root.querySelector("[data-ac-datepicker-trigger]");
    hidden.value = iso;
    trigger.value = display;
    dpClose(root);
  }

  // Event delegation for datepicker clicks
  document.addEventListener("click", function (e) {
    var root;

    // Trigger opens picker
    if (e.target.closest("[data-ac-datepicker-trigger]")) {
      root = e.target.closest("[data-ac-datepicker]");
      if (root) dpOpen(root);
      return;
    }

    // Close button
    if (e.target.closest("[data-ac-datepicker-close]")) {
      root = e.target.closest("[data-ac-datepicker]");
      if (root) dpClose(root);
      return;
    }

    // Cancel button
    if (e.target.closest("[data-ac-datepicker-cancel]")) {
      root = e.target.closest("[data-ac-datepicker]");
      if (root) dpClose(root);
      return;
    }

    // Overlay background click
    if (e.target.classList.contains("ac-datepicker-overlay")) {
      root = e.target.closest("[data-ac-datepicker]");
      if (root) dpClose(root);
      return;
    }

    // Confirm button
    if (e.target.closest("[data-ac-datepicker-confirm]")) {
      root = e.target.closest("[data-ac-datepicker]");
      if (root) dpConfirm(root);
      return;
    }

    // Today button
    if (e.target.closest("[data-ac-datepicker-today]")) {
      root = e.target.closest("[data-ac-datepicker]");
      if (root) {
        var state = dpGetState(root);
        var now = new Date();
        state.selYear = now.getFullYear();
        state.selMonth = now.getMonth();
        state.selDay = now.getDate();
        state.viewYear = now.getFullYear();
        state.viewMonth = now.getMonth();
        state.step = "day";
        dpRender(root);
      }
      return;
    }

    // Back button
    if (e.target.closest("[data-ac-datepicker-back]")) {
      root = e.target.closest("[data-ac-datepicker]");
      if (root) {
        var s = dpGetState(root);
        if (s.step === "day") s.step = "month";
        else if (s.step === "month") s.step = "year";
        dpRender(root);
      }
      return;
    }

    // Year cell
    var yearCell = e.target.closest("[data-ac-datepicker-year]");
    if (yearCell) {
      root = yearCell.closest("[data-ac-datepicker]");
      if (root) {
        var st = dpGetState(root);
        st.viewYear = parseInt(yearCell.dataset.acDatepickerYear, 10);
        st.selYear = st.viewYear;
        st.step = "month";
        dpRender(root);
      }
      return;
    }

    // Month cell
    var monthCell = e.target.closest("[data-ac-datepicker-month]");
    if (monthCell) {
      root = monthCell.closest("[data-ac-datepicker]");
      if (root) {
        var sm = dpGetState(root);
        sm.viewMonth = parseInt(monthCell.dataset.acDatepickerMonth, 10);
        sm.selMonth = sm.viewMonth;
        sm.step = "day";
        dpRender(root);
      }
      return;
    }

    // Day cell
    var dayCell = e.target.closest("[data-ac-datepicker-day]");
    if (dayCell) {
      root = dayCell.closest("[data-ac-datepicker]");
      if (root) {
        var sd = dpGetState(root);
        sd.selDay = parseInt(dayCell.dataset.acDatepickerDay, 10);
        dpRender(root);
      }
      return;
    }
  });

  // Double-click day to instant confirm
  document.addEventListener("dblclick", function (e) {
    var dayCell = e.target.closest("[data-ac-datepicker-day]");
    if (dayCell) {
      var root = dayCell.closest("[data-ac-datepicker]");
      if (root) {
        var state = dpGetState(root);
        state.selDay = parseInt(dayCell.dataset.acDatepickerDay, 10);
        dpConfirm(root);
      }
    }
  });

  // Escape closes open datepickers
  document.addEventListener("keydown", function (e) {
    if (e.key === "Escape") {
      var openPickers = document.querySelectorAll(
        '.ac-datepicker-overlay[style*="flex"]'
      );
      openPickers.forEach(function (overlay) {
        var root = overlay.closest("[data-ac-datepicker]");
        if (root) dpClose(root);
      });
    }
  });

  // Global helpers
  window.acOpenDatePicker = function (id) {
    var root = document.getElementById(id);
    if (root && root.hasAttribute("data-ac-datepicker")) dpOpen(root);
  };

  window.acCloseDatePicker = function (id) {
    var root = document.getElementById(id);
    if (root && root.hasAttribute("data-ac-datepicker")) dpClose(root);
  };
})();
