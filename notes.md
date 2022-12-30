|                       | normal clone | bare clone     | worktree repo              |
| --------------------- | ------------ | -------------- | -------------------------- |
| <DIR>                 | /tmp/prompt  | /tmp/prompt_wt | /tmp/prompt_wt/master      |
| git-dir               | .git         | .              | prompt_wt/worktrees/master |
| absolute-git-dir      | prompt/.git  | prompt_wt      | prompt_wt/worktrees/master |
| show-toplevel         | prompt       |                | prompt_wt/master           |
| git-common-dir        | .git         | .              | prompt_wt                  |
| is-bare-repository    | FALSE        | TRUE           | FALSE                      |
| is-shallow-repository | FALSE        | FALSE          | FALSE                      |
| is-inside-git-dir     | FALSE        | TRUE           | FALSE                      |
| is-inside-work-tree   | TRUE         | FALSE          | TRUE                       |
