# d2-lsp

Creating language server protocol(lsp) for [D2lang](https://d2lang.com/).
I am creating this after learning about [D2lang](https://d2lang.com/) and contributing to
its [tree-sitter](https://github.com/ravsii/tree-sitter-d2)

## Building

To build the lsp to test locally, please build the project to generate *main* binary.

```sh
go build main.go
```

## Neovim Lsp Setup

To setup and test this Lsp locally for *.d2* files please add this to your nvim configuration

```lua
local client = vim.lsp.start_client({
  name = "D2lsp",
  cmd = { "/path/to/d2-lsp/main" },
})

if not client then
  vim.notify("Lsp not well setup")
end

vim.api.nvim_create_autocmd("FileType", {
  pattern = "d2",
  callback = function()
    vim.lsp.buf_attach_client(0, client)
  end,
})
```

## Reference

1. [Microsoft Specification](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/)
