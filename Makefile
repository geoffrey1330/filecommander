# https://github.com/geoffrey1330/filecommander

BIN_NAME=filecommander

BIN_ROOT=$(PWD)/.bin
export PATH:=$(PATH):$(BIN_ROOT)

print:

bin-init:
	rm -rf $(BIN_ROOT)
	mkdir -p $(BIN_ROOT)
bin: bin-init
	go build -o $(BIN_ROOT)/$(BIN_NAME) .

### run and tests

RUN_ROOT=$(PWD)/.test
RUN_CMD=cd $(RUN_ROOT) && $(BIN_NAME)

run-init-del:
	rm -rf $(RUN_ROOT)
run-init: run-init-del
	mkdir -p $(RUN_ROOT)

run-h:
	$(RUN_CMD) -h

run-create: run-init
	$(RUN_CMD) create test.md
	# fails: does not do folders
	$(RUN_CMD) create testfolder
	$(RUN_CMD) write test.md bbb-ddd
	$(RUN_CMD) read test.md
	$(RUN_CMD) list .



run-delete-pre: run-init
	touch $(RUN_ROOT)/delete.md
	touch $(RUN_ROOT)/delete01.md
	touch $(RUN_ROOT)/delete01.text
	mkdir -p  $(RUN_ROOT)/folder-delete01
	mkdir -p  $(RUN_ROOT)/folder-delete02
	touch $(RUN_ROOT)/folder-delete02/delete01.text
run-delete-h:
	# fail: No help.
	$(RUN_CMD) delete -h
run-delete: run-delete-pre
	# fail: No help.
	$(RUN_CMD) delete -h
	# pass
	$(RUN_CMD) delete delete.md
	# pass
	$(RUN_CMD) delete delete*.md
	# pass
	$(RUN_CMD) delete delete*.*
	# pass. empty folder
	$(RUN_CMD) delete folder-delete01
	# pass. non empty folder. Need "-f" flag...
	$(RUN_CMD) delete folder-delete02
	
	
run-copy-pre: run-init
	touch $(RUN_ROOT)/specific-copy.md
	touch $(RUN_ROOT)/extension-copy.md
	touch $(RUN_ROOT)/extension-copy.txt
	touch $(RUN_ROOT)/filename-copy01.text
	touch $(RUN_ROOT)/filename-copy02.text
run-copy-h:
	# fail: No help.
	$(RUN_CMD) copy -h
run-copy: run-copy-pre
	# pass: specific
	$(RUN_CMD) copy specific-copy.md end-specific-copy.md
	# fail: extension. Make it .txt. weird ?
	$(RUN_CMD) copy extension-copy.* end-extension-copy.md
	# fail: filename wildcard. 
	$(RUN_CMD) copy filename-copy*.* end-filename-copy.*

run-move-pre: run-init
	touch $(RUN_ROOT)/specific-move.md
	touch $(RUN_ROOT)/extension-move.md
	touch $(RUN_ROOT)/extension-move.txt
	touch $(RUN_ROOT)/filename-move01.text
	touch $(RUN_ROOT)/filename-move02.text
run-move: run-move-pre
	# fail: No wildcards.
	$(RUN_CMD) move -h

	# pass: specific
	$(RUN_CMD) move specific-move.md end-specific-move.md

	# fail: extension wildcard
	#$(RUN_CMD) move extension-move.* env-extension-move/*

	# fail: filename wildcard
	$(RUN_CMD) move filename-e*.* end-filename-move/*

run-search-pre:
	mkdir -p $(RUN_ROOT)/search
	touch $(RUN_ROOT)/search/test01.md
run-search-h:
	# fails
	$(RUN_CMD) search -h
run-search: run-search-pre
	
	mkdir -p ./.test

	@echo ""
	@echo "Can find by wildcard file extension ? YES"
	$(RUN_CMD) search ./search/test1 test01.*
	@echo ""

	@echo ""
	@echo "Cant find by wildcard withng file name ? NO"
	mkdir -p ./.test/test2
	touch ./.test/test2/aaa_bin_darwin_arm64
	$(RUN_CMD) search ./.test/test2 *_bin_*
	@echo ""

    # Cant find by size ?
	@echo ""
	@echo "Cant find by size ? NO"
	mkdir -p ./.test/test2
	touch ./.test/test2/bigfile.md
	# ex: # find . -type f -size +20M
	$(RUN_CMD) search ./.test/test2 -size +20M
	@echo ""

	
run-open:
	# None work. Its calling cmd.exe("editor").
	# We probably need a "open native" command too, which allows the OS's own scheme stuff to pick what to do.
	$(RUN_CMD) open -h
	$(RUN_CMD) open /Applications/Chapar.app
	$(RUN_CMD) open test.md
	

	
	
	



	
	

