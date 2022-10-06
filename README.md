# downloader

Downloader is CLI library to download media from multiple platform. For now, it only supports video or audio from YouTube and videos from Tiktok.

## Prerequisite

- FFMPEG

    ffmpeg is used to combining video and audio files from youtube.

    You can download ffmpeg [here](https://ffmpeg.org/download.html).

    Or on linux:

    ```bash
    sudo apt install ffmpeg
    ```

## Binary Download

You can download binary files in the [Release Page](https://github.com/portalnesia/downloader/releases).


----


## CLI Usage

- Go to downloader binary folder

    ```bash
    cd path/to/downloader
    ```

- To show help messages about cli usage. You can type bellow command in your terminal/cmd

    Linux:

    ```bash
    ./downloader -h
    ```

    Windows:

    ```bash
    downloader -h
    ```

### Youtube

```bash
./downloader youtube [-h] [-i] [-o] [url youtube]
```

| Flags | Alias | Description |
| ---- | --- | --- |
| `--help` | `-h` | Get help messages for youtube commands |
| `--info` | `-i` | Get youtube video information, without downloading video or audio |
| `--output` | `-o` | Set output directory |

### Tiktok

```bash
./downloader tiktok [-h] [-i] [-o] [url tiktok]
```

| Flags | Alias | Description |
| ---- | --- | --- |
| `--help` | `-h` | Get help messages for tiktok commands |
| `--info` | `-i` | Get tiktok video information, without downloading video |
| `--output` | `-o` | Set output directory |

