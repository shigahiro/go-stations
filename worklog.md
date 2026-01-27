以下の手順でコンフリクトを解消した。

# 1. フォーク元をリモートとして追加（まだの場合）

git remote add upstream https://github.com/TechBowl-japan/go-stations.git

# 2. フォーク元の最新情報を取得

git fetch upstream

# 3. フォーク元のmainブランチを自分のブランチにマージ

git merge upstream/main
