```
This project is under active development and is not ready for production use.
```

# Infinite Ez

Infinite Ez is a proprietary container management platform in a single file. It has a REST API, CLI and dashboard.

## Running

## Development

In this repository you'll find the REST API and CLI code plus the dashboard assets. The API and CLI uses Clean Architecture, DDD, TDD, CQRS, Object Calisthenics, etc. Understand how these concepts works before proceeding is advised.

To run this project during development you must install [Air](https://github.com/cosmtrek/air). Air is a tool that will watch for changes in the project and recompile it automatically.

### Unit Testing

Since Ez relies on the operational system being openSUSE MicroOS, the entire development and testing should be done in a VM. The VM can be created with the following steps:

1. Install VMWare Player;

2. Download the VMware `.vmx` and `.vmdk` files from "Base System + Container Runtime" column on MicroOS download page:
   https://en.opensuse.org/Portal:MicroOS/Downloads

Note: make sure you downloaded the "Base System + Container Runtime" files and not the "Base System" files.

3. Extract the `.vmdk` file if it came compressed.

4. Add the VM to the VMWare Player interface using the "Open a Virtual Machine" option and then:

   1. Change the resources to 4GB RAM // 2vCPU;
   2. Remove the default disk;
   3. Remove the floppy drive;
   4. Add a new disk selecting "Use an existent disk" and use the .vmdk file you downloaded (keep the format when asked);
   5. Add a secondary disk (5GB minimal - do not format);
   6. Add a Network Adapter (keep set to NAT mode);

5. Run the VM. The first boot will allow you to set up a root password. Keep the keyboard to US layout to avoid issues. After the first reboot, login with the password you set.

6. Add your public SSH key to `/root/.ssh/authorized_keys` file. I would recommend uploading your .pub to a Pastebin, File.io or similar service and then using curl to download it to the VM as curl is installed by default.

Note: If you use Pastebin, make sure you're downloading the raw file and not the HTML page.

7. Install git and Go and reboot:

```
transactional-update pkg install git make
curl -L https://go.dev/dl/go1.23.1.linux-amd64.tar.gz -o go.tar.gz
tar -C /usr/local -xzf go.tar.gz
rm -f go.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin:~/go/bin' >> ~/.bashrc
echo 'alias ez-swag="swag init -g src/presentation/api/api.go -o src/presentation/api/docs"' >> ~/.bashrc
systemctl reboot
```

8. After the reboot, you'll need to configure GitHub authentication. The `.ssh/config` file is your friend, this is an example:

```
Host github.com
  HostName github.com
  User git
  IdentityFile ~/.ssh/the_github_key
```

Replace `the_github_key` with the path to your private key and remember to chmod the key to 400.

9. Install a few Go packages and clone the Ez repository:

```
go install github.com/swaggo/swag/cmd/swag@latest
go install github.com/cosmtrek/air@latest
go install github.com/a-h/templ/cmd/templ@latest
git config --global user.name "yourgithubnick"
git config --global user.email yourgithubemail
git clone git@github.com:goinfinite/ez.git
```

10. Build the project and run the installer:

```
cd ez; air # hit CTRL-C after running this
chmod +x /var/infinite/ez
/var/infinite/ez sys-install
```

11. **The system will reboot** and once see Ez running via systemd, you should be able to use the [Visual Studio Remote SSH extension](https://code.visualstudio.com/docs/remote/ssh) to connect to the VM and manage the project.

Make sure to use the SSH key to connect to the VM and not the password. The IP address of the VM can be found on the VMWare interface which you used to SSH. If you don't see it, just use `exit` to return to the login screen and the IP will be displayed there.

You should also install these VSCode extensions during development:

```
EditorConfig.EditorConfig
GitHub.copilot
GitHub.vscode-pull-request-github
esbenp.prettier-vscode
foxundermoon.shell-format
golang.go
hbenl.vscode-test-explorer
ms-vscode.test-adapter-converter
redhat.vscode-yaml
streetsidesoftware.code-spell-checker
streetsidesoftware.code-spell-checker-portuguese-brazilian
timonwong.shellcheck
bradlc.vscode-tailwindcss
Trapfether.tailwind-raw-reorder
a-h.templ
```

12. The install process also creates a systemd unit to run the project. Use the following commands to disable the service so you can run the project manually with Air during development:

```
systemctl disable ez
systemctl stop ez
```

13. If you get a license error, you can run `ez license refresh` to force the license to be regenerated.

### Environment Variables

You must have an `.env` file in the root of the git directory **during development**. You can use the `.env.example` file as a template. Air will read the `.env` file and use it to run the project during development.

If you add a new env var that is required to run the apis, please add it to the `src/presentation/shared/checkEnvs.go` file.

When running in production, the `/var/infinite/.env` file is only used if the environment variables weren't set in the system. For instance, if you want to set the `ENV1` variable, you can do it in the `.env` file or in the command line:

```
ENV1=XXX /var/infinite/ez
```

### Dev Utils

The `src/devUtils` folder is not a Clean Architecture layer, it's there to help you during development. You can add any file you want there, but it's not recommended to add any file that is not related to development since the code there is meant to be ignored by the build process.

For instance there you'll find a `testHelpers.go` file that is used to read the `.env` during tests.

### Building

Ez is likely on the marketplace of your cloud provider already, but if you want to build it yourself.

The software itself is a single binary, but it requires openSUSE MicroOS to run.

1. Once you have uploaded the openSUSE MicroOS cloud-init image to your provider, attach a secondary unformatted disk and deploy the VM.

_If you are running locally, follow steps 1 to 6 of the "Unit Testing" title above._

2. Get the Ez binary, download it to the `/var/infinite/` directory and give it execution permission:

```
chmod +x /var/infinite/ez
```

3. Run the installer:

```
/var/infinite/ez sys-install
```

4. The system will reboot and once you see Ez binary running via systemd, you should be able to access the dashboard on the VM IP address on port 3141.

There will be no log messages on the console for now, the installation process will be refactored soon so that the entire process happens on the Web UI.

### Web UIs

This project has two web UIs, the previous Vue.js frontend and the new HTMX frontend. The Vue.js frontend is deprecated and will be removed in the future. It's available at `/_/` and the HTMX frontend is available at `/`.

The HTMX frontend has a hot reload feature that will reload the page when the assets are changed. It's based on a websocket connection that the client listens to and will reload the page when the server stops responding (which is when Air is rebuilding the binary after a file change). To enable this feature, simply put a `DEV_MODE=true` in the `.env` file.

## REST API

The REST API is exposed on port 3141 via HTTPS on the endpoint `/api/`. The API endpoint are versioned and the version is part of the URL.

### Swagger // OpenAPI

The API is documented using Swagger/OpenAPI 2.0 and the documentation can be found on the `/api/swagger/` endpoint.

To generate the swagger documentation, you must use the following command:

```
swag init -g src/presentation/api/api.go -o src/presentation/api/docs
```

The annotations are in the controller files. The reference file can be found [here](https://github.com/swaggo/swag#attribute).

### Authentication

The API accepts two types of tokens and uses the standard "Authorization: Bearer \<token\>" header:

- **sessionToken**: is a JWT, used for dashboard access and generated with the user login credentials. The token contains the accountId, IP address and expiration date. It expires in 3 hours and only the IP address used on the token generation is allowed to use it.

- **accountApiKey**: is a token meant for M2M communication. The token is a _AES-256-CTR-Encrypted-Base64-Encoded_ string, but only the SHA3-256 hash of the key is stored in the server. The accountId is retrieved during key decoding, thus you don't need to provide it. The token never expires, but the user can update it at any time.
