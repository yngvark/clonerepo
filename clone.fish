function clone --wraps='/home/yngvar/.go/bin/clonerepo ' --wraps='cd $(/home/yngvar/.go/bin/clonerepo )' --description 'alias clone=cd $(/home/yngvar/.go/bin/clonerepo )'
  #echo $argv[1]

  switch $argv[1]
      case "git@*"
        cd $(clonerepo $argv)
      case "https*"
        cd $(clonerepo $argv)
      case '*'
        clonerepo $argv
  end
end
