

document.addEventListener("DOMContentLoaded", function () {
    initializeYearPicker();
});

function toggleDropdown() {
    const yearDropdown = document.getElementById('yearDropdown');
    if (yearDropdown) {
        yearDropdown.classList.toggle('hidden');
    }
}

document.body.addEventListener('htmx:afterSwap', (event) => {
    
    initializeYearPicker();
});
function initializeYearPicker() {
    const yearDropdown = document.getElementById('yearDropdown');
    const yearInput = document.getElementById('yearInput');

    if (!yearDropdown || !yearInput) {
        console.warn("Year picker elements not found. Initialization deferred.");
        return;
    }

    const currentYear = new Date().getFullYear();
    const startYear = currentYear - 10;
    const endYear = currentYear + 10;

    yearDropdown.innerHTML = '';

    for (let year = startYear; year <= endYear; year++) {
        const yearOption = document.createElement('div');
        yearOption.className = "p-2 cursor-pointer hover:bg-blue-100";
        yearOption.textContent = year;
        yearOption.setAttribute("hx-get", `/data?time=${year}`);
        yearOption.setAttribute("hx-trigger", "click");
        yearOption.setAttribute("hx-target", "#content");
        yearOption.setAttribute("hx-swap", "#innerHTML");
        yearOption.onclick = () => selectYear(year);
        yearDropdown.appendChild(yearOption);
    }

    // Process HTMX for dynamically added elements
    htmx.process(yearDropdown);

    function toggleDropdown() {
        yearDropdown.classList.toggle('hidden');
    }

    yearInput.addEventListener("click", toggleDropdown);

    function selectYear(year) {
        yearInput.value = year;
        yearDropdown.classList.add('hidden');
    }

    window.addEventListener('click', (event) => {
        if (!event.target.closest('.relative')) {
            yearDropdown.classList.add('hidden');
        }
    });
}
function initializeFlatpickr(target) {
            flatpickr(target, {
                dateFormat: "d-m-Y", // Custom date format
                onChange: function(selectedDates, dateStr, instance) {
                    // Trigger HTMX manually when a date is selected
                    target.dispatchEvent(new Event('change', { bubbles: true }));
                }
            });
            console.log("Flatpickr initialized on:", target);
        }

        // Initialize Flatpickr on page load
        document.addEventListener('DOMContentLoaded', function () {
            const input = document.querySelector('#datepicker-format');
            if (input) {
                initializeFlatpickr(input);
            }
        });

        // Reinitialize Flatpickr after HTMX content updates
        document.body.addEventListener('htmx:afterRequest', function (event) {
            // Check if the swapped content contains the date-picker input
            const newInput = event.target.querySelector('#datepicker-format');
            if (newInput) {
                initializeFlatpickr(newInput);
            }
        });
         document.body.addEventListener('htmx:beforeRequest', function (event) {
            // Check if the swapped content contains the date-picker input
            const newInput = event.target.querySelector('#datepicker-format');
            if (newInput) {
                initializeFlatpickr(newInput);
            }
        });
