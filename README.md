# os-scrapper

OS information scrapper agents

### TO-DO

- [ ] Distro Information

  Distribution name, version, and codename.
  Kernel version and build information.
  Bootloader information (e.g., GRUB settings).

- [ ] Directory Structure

  Overview of the root directory structure.
  Size and usage statistics of major directories.
  Mounted file systems and their mount points.

- [ ] Shell Information

  Default shell and available shells.
  Shell configuration files (.bashrc, .zshrc, etc.).
  Shell aliases and functions.

- [ ] Environment Information

  System-wide environment variables.
  User-specific environment variables.
  Timezone and locale settings.
  System time and date.

- [ ] Language-Based Environment Information

  Python environment details (venv, pyenv, etc.).
  Node.js environment (nvm, .node-version, etc.).
  Ruby environment (RVM, rbenv, etc.).
  Java version and related settings.
  Go environment variables (GOPATH, GOROOT, etc.).
  PHP environment details.

- [ ] Installed Packages for Python, Node.js (Both Global and Local)

  List of globally installed Python packages (pip list).
  List of Python packages in virtual environments.
  List of globally installed Node.js packages (npm -g list).
  List of locally installed Node.js packages in projects.

- [ ] Wi-Fi Information, Bluetooth Information (Only Hardware Information, Nothing Sensitive)

  Details about Wi-Fi hardware (chipset, driver, etc.).
  Bluetooth hardware details (chipset, driver, etc.).
  Status of Wi-Fi and Bluetooth interfaces. Network connectivity

- [ ] Hardware CPU-Based Information

  CPU model, architecture, and core details.
  CPU flags and features.
  Temperature sensors and power management settings.

- [ ] Process Information

  List of running processes.
  Process tree structure.
  Top resource-consuming processes.
  Details of specific critical processes (e.g., init, sshd).

- [ ] Storage, of All Types

  Disk usage and partition information (df, fdisk).
  Details of mounted file systems.
  RAID configuration details.
  LVM (Logical Volume Management) details.
  Swap space usage.

- [ ] Installed Programs and Available Packages

  List of installed packages and versions.
  List of available packages from configured repositories.
  Flatpak, Snap, and AppImage package details (if applicable).
  Locally compiled or manually installed software.

- [ ] Repo Information and Cache

  Information on configured package repositories (apt, yum, dnf, etc.).
  Cache sizes and contents for package managers.
  Repo keys and GPG signatures.

- [ ] Cron Jobs and Other Scheduled Tasks

  User-specific cron jobs (crontab -l).
  System-wide cron jobs (/etc/crontab, /etc/cron.d/).
  Systemd timers and scheduled tasks.
  Anacron, at jobs, and other task scheduling configurations.

- [ ] Kernel Modules

  List of loaded kernel modules.
  Details about available kernel modules.
  Configuration files for kernel modules.

- [ ] Network Configuration

  Active network interfaces and their configurations.
  IP addresses, MAC addresses, and routing tables.
  DNS server configurations.
  Network connections and listening ports (netstat, ss).

- [ ] Firewall and Security Information

  Firewall rules (iptables, firewalld, ufw).
  SELinux or AppArmor status and configurations.
  Installed security tools (e.g., fail2ban, auditd).

- [ ] Running Services and Daemons

  List of active services (systemctl list-units --type=service).
  Startup settings for services (enabled, disabled).
  Daemon-specific configurations (e.g., sshd_config).

- [ ] Logs and Journals

  System logs (/var/log/).
  Recent journal entries (journalctl).
  Error and warning logs.

- [ ] User and Group Information

  List of users and groups (/etc/passwd, /etc/group).
  User home directories and shell settings.
  User login history and session details (last, w).

- [ ] Virtualization Information

  Check if running in a virtualized environment.
  Details about the hypervisor (KVM, VMware, etc.).
  Virtual machine configurations (if applicable).

- [ ] Power and Battery Information

  Battery status and health (for laptops).
  Power management settings and configurations.
  ACPI information and related logs.

- [ ] Available and Loaded Kernel Modules

  List of available kernel modules.
  Details about currently loaded modules.

- [ ] System Uptime and Load

  System uptime.
  Current load average.
  CPU and memory usage statistics (top, htop).

#### ... Add more

## Rules

- DONT PUSH TO MAIN
- All agents should be under agent/agentname
- Make sure to pull or fetching before working

## Contributors

- [Athul Prakash NJ](https://github.com/psychoSherlock)
- [Andrew C Anil](https://github.com/iamandrewcanil)
- [Akshhay KM](https://github.com/Xanthium7)
- [Abin Joy](https://github.com/Abinjoy025)
- [Anubind C Biju](https://github.com/anubix05)
- [Vaishakh S Nair (OneGrit)](https://github.com/vaishakhsnair)

## Workflow

1. Create a Branch for Each Feature or Bug Fix

   When you start working on a new feature or bug fix, create a new branch from the main or develop branch. The branch should have a descriptive name that indicates the purpose, such as feature/add-user-auth or bugfix/fix-login-issue.
   This keeps your work separate from the stable code in the main branch, reducing the risk of introducing issues.

2. Work on the Feature or Fix in Your Branch

   Make all your changes in the feature branch. This includes writing code, adding tests, and making any necessary documentation updates.
   By isolating your work in a branch, you can experiment and make multiple commits without affecting the main codebase.

3. Push the Branch to the Main Repository

   Once you've made some progress or completed your work, push the branch to the main repository on GitHub. This makes your work visible to the rest of the team and ensures it's backed up remotely.
   Regularly pushing your branch helps you collaborate with others and allows teammates to review your work early if needed.

4. Create a Pull Request (PR)

   After you've completed the feature or fix, create a Pull Request (PR) from your feature branch to the main or develop branch.
   The PR is a formal request to merge your changes into the main codebase. It allows the team to review your work, discuss any potential issues, and ensure that the feature or fix meets the project's standards.

5. Review and Discuss the PR

   Team members review the PR, leaving comments or suggestions for improvements. This is an opportunity to catch bugs, improve code quality, and ensure that the changes align with the overall project goals.
   If changes are requested, you can update your branch with additional commits to address the feedback.

The Feature Branch Workflow allows us to work on new features or fixes in isolation, reducing the risk of breaking the main codebase. Each feature branch is a safe space to develop and test changes without affecting others. By pushing branches and creating PRs, we maintain a clear process for reviewing and integrating changes, ensuring that our code remains stable and high-quality.

This workflow supports collaboration by enabling multiple developers to work on different features simultaneously. It also provides a clear history of changes, making it easier to track the progress and evolution of the project.
