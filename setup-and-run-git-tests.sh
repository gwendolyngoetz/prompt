#! /bin/bash

DEBUG="${1}"

ROOT_FOLDER="/tmp/testing"
PROMPT_FOLDER="${ROOT_FOLDER}/prompt"
PROMPT_WT_FOLDER="${ROOT_FOLDER}/prompt_wt"

mkdir -p ${ROOT_FOLDER}

# Compile git tests
cd pkg/git
go test -c
mv git.test ${ROOT_FOLDER}

# Clean
rm -rf ${PROMPT_FOLDER}
rm -rf ${PROMPT_WT_FOLDER}

# Setup
mkdir -p ${PROMPT_FOLDER}
cd ${PROMPT_FOLDER}
git init

# Add file1
FILE1='file1.txt'
echo "test" >> ${FILE1}
git add ${FILE1}
git commit -m "Adding ${FILE1}"

# Add file2
FILE2='file2.txt'
echo "test" >> ${FILE2}
git add ${FILE2}
git stash push ${FILE2} -m "Adding ${FILE2}"

# Add pkg test folder
mkdir -p ${PROMPT_FOLDER}/pkg/git
cd ${PROMPT_FOLDER}/pkg/git

FILE3='file3.txt'
echo "test" >> ${FILE3}
git add ${FILE3}
git commit -m "Adding ${FILE3}"

# Add branches
git branch branch1
git branch branch2

# Clone bare
cd ${PROMPT_FOLDER}
cd ..
git clone --bare prompt prompt_wt
cd prompt_wt

# Add worktree branches
git worktree add master
git worktree add branch1
git worktree add branch2
cd master

# Add file4
FILE4='file4.txt'
echo "test" >> ${FILE4}
git add ${FILE4}
git stash push ${FILE4} -m "Adding ${FILE4}"

# Copy git.test
cp ${ROOT_FOLDER}/git.test ${PROMPT_FOLDER}/pkg/git
cp ${ROOT_FOLDER}/git.test ${PROMPT_WT_FOLDER}/master/pkg/git

# Clear 
if [[ -z "${DEBUG}" ]]; then
    clear
fi

# Run git.test
echo "PROMPT FOLDER"
cd ${PROMPT_FOLDER}/pkg/git
./git.test

echo "PROMPT WORKTREE FOLDER"
cd ${PROMPT_WT_FOLDER}/master/pkg/git
./git.test

