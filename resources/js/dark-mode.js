const themeToggleDarkIcon = document.getElementById('theme-toggle-dark-icon');
const themeToggleLightIcon = document.getElementById('theme-toggle-light-icon');

// change the icons inside the button based on previous settings
if (
  localStorage.getItem('color-theme') === 'light' ||
  ('color-theme' in localStorage &&
    localStorage.getItem('color-theme') !== 'dark') ||
  (!('color-theme' in localStorage) &&
    !window.matchMedia('(prefers-color-scheme: light)').matches)
) {
  document.documentElement.classList.remove('dark');
  themeToggleDarkIcon?.classList.remove('hidden');
} else {
  document.documentElement.classList.add('dark');
  themeToggleLightIcon?.classList.remove('hidden');
}

const event = new Event('dark-mode');
const themeToggleBtn = document.getElementById('theme-toggle');

themeToggleBtn?.addEventListener('click', function () {
  // toggle icons
  themeToggleDarkIcon?.classList.toggle('hidden');
  themeToggleLightIcon?.classList.toggle('hidden');

  if (localStorage.getItem('color-theme')) {
    // if set via local storage previously
    if (localStorage.getItem('color-theme') === 'light') {
      document.documentElement.classList.add('dark');
      localStorage.setItem('color-theme', 'dark');
    } else {
      document.documentElement.classList.remove('dark');
      localStorage.setItem('color-theme', 'light');
    }
  } else {
    // if NOT set via local storage previously
    if (document.documentElement.classList.contains('dark')) {
      document.documentElement.classList.remove('dark');
      localStorage.setItem('color-theme', 'light');
    } else {
      document.documentElement.classList.add('dark');
      localStorage.setItem('color-theme', 'dark');
    }
  }

  document.dispatchEvent(event);
});
