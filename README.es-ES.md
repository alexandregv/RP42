

# RP42

Una integración de Discord Rich Presence para [@42School](https://github.com/42School).  
![Captura de pantalla](https://raw.githubusercontent.com/alexandregv/RP42/refs/heads/master/screenshot.png)

<!--TOC-->

- [Instalación](#installation)
- [Uso](#usage)
- [Compilarlo tú mismo](#building-yourself)
- [Contribuir](#contributing)
- [Colaboradores](#contributors)
- [Estrellas a lo largo del tiempo](#stargazers-over-time)

<!--TOC-->

## Instalación

1. Descarga RP42 desde la página de [lanzamientos](https://github.com/alexandregv/RP42/releases), o compílalo tú mismo
2. Crea una aplicación de API en la Intranet: https://profile.intra.42.fr/oauth/applications/new

Alternativamente, si tienes Go instalado:

```bash
go install github.com/alexandregv/RP42@latest
```

/!\ NO compartas tus credenciales de API con otra persona, ni en GitHub, etc. /!\

## Uso

Ejecuta la aplicación de la siguiente manera, utilizando las credenciales de tu aplicación de API:

- Linux: `./RP42 -i CLIENT_ID -s CLIENT_SECRET &`
- Windows: `RP42.exe -i CLIENT_ID -s CLIENT_SECRET`
- MacOS: `./RP42 -i CLIENT_ID -s CLIENT_SECRET`

En Linux, para ejecutarlo en segundo plano como un servicio, crea un archivo `~/.config/systemd/user/RP42.service` con este contenido:

```ini
[Unit]
Description=Discord Rich Presence para 42

[Service]
ExecStart=%h/.local/bin/RP42 -i <my-app-id> -s <my-app-secret>
Restart=always

[Install]
WantedBy=default.target
```

Luego ejecuta `systemctl --user start RP42` y `systemctl --user status RP42`.

## Compilarlo tú mismo

Si deseas compilar RP42 tú mismo, sigue estas instrucciones:

1. Clona el repositorio: `git clone https://github.com/alexandregv/RP42.git`
2. Compila: `make` para todas las distribuciones, o `make linux` / `make macos` / `make windows`

## Contribuir

1. Haz un fork (<https://github.com/alexandregv/RP42/fork>)
2. Crea tu rama de características (`git checkout -b my-new-feature`)
3. Confirma tus cambios (`git commit -am 'Add some feature'`)
4. Envía a la rama (`git push origin my-new-feature`)
5. Crea un nuevo Pull Request

## Colaboradores

- [alexandregv/aguiot--](https://github.com/alexandregv) - creador y mantenedor

## Estrellas a lo largo del tiempo

[![Estrellas a lo largo del tiempo](https://starchart.cc/alexandregv/RP42.svg?variant=adaptive)](https://starchart.cc/alexandregv/RP42)
