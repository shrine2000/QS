# QuickStart

QuickStart accelerates project setup by automating folder creation, Git initialization, and README/LICENSE generation based on user-defined parameters. It streamlines the process, ensuring a consistent and well-documented project structure. Finally, it opens the project in Visual Studio Code.

To build the CLI executable, run the following command in the project directory:

```shell
go build -o qs
```

This will create an executable file named "qs". Move this executable to a directory that's in your system's PATH, such as /usr/local/bin or $HOME/go/bin, to use the "qs" command globally.

```shell
sudo mv qs /usr/local/bin
```

Now, `qs` command can be used as 

```shell
qs -name my-app -path /path/to/directory -project "My App"
```


