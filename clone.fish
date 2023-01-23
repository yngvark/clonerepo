function clone --wraps='/home/yngvar/.go/bin/clonerepo ' --wraps='cd $(/home/yngvar/.go/bin/clonerepo )' --description 'alias clone=cd $(/home/yngvar/.go/bin/clonerepo )'
  set -l OUT $(clonerepo $argv)

  if test (count $OUT) -eq 1
    cd $(clonerepo $argv)
  end
end
