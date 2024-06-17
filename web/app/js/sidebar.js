const sidebar = document.getElementById('sidebar');

const toggleSidebarMobile = (
  sb,
  sidebarBackdrop = null,
  toggleSidebarMobileHamburger = null,
  toggleSidebarMobileClose = null,
) => {
  sb.classList.toggle('hidden');
  sidebarBackdrop?.classList.toggle('hidden');
  toggleSidebarMobileHamburger?.classList.toggle('hidden');
  toggleSidebarMobileClose?.classList.toggle('hidden');
};

if (sidebar) {
  const toggleSidebarMobileEl = document.getElementById('toggleSidebarMobile');
  const sidebarBackdrop = document.getElementById('sidebarBackdrop');
  const toggleSidebarMobileHamburger = document.getElementById(
    'toggleSidebarMobileHamburger',
  );
  const toggleSidebarMobileClose = document.getElementById(
    'toggleSidebarMobileClose',
  );
  const toggleSidebarMobileSearch = document.getElementById(
    'toggleSidebarMobileSearch',
  );

  toggleSidebarMobileSearch?.addEventListener('click', () => {
    sidebar.classList.add('flex');
    toggleSidebarMobile(
      sidebar,
      sidebarBackdrop,
      toggleSidebarMobileHamburger,
      toggleSidebarMobileClose,
    );
  });

  toggleSidebarMobileEl?.addEventListener('click', () => {
    sidebar.classList.add('flex');
    toggleSidebarMobile(
      sidebar,
      sidebarBackdrop,
      toggleSidebarMobileHamburger,
      toggleSidebarMobileClose,
    );
  });

  sidebarBackdrop?.addEventListener('click', () => {
    toggleSidebarMobile(
      sidebar,
      sidebarBackdrop,
      toggleSidebarMobileHamburger,
      toggleSidebarMobileClose,
    );
  });
}
