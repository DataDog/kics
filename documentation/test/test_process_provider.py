import pytest
import local_gen

from unittest.mock import mock_open, patch, MagicMock
from pathlib import Path

class TestProcessProvider:
    @pytest.fixture
    def test_path(self):
        return Path("/")
    
    @pytest.fixture
    def test_string(self):
        return "aws"

    def test_provider_is_not_directory(self, test_path):
        with patch.object(Path, 'is_dir', return_value=False) as mock_is_dir:
            result = local_gen.process_provider(test_path, test_path, test_path, test_path, 0, [], {})
            assert result == 0
            mock_is_dir.assert_called_once()
    
    def test_metadata_failure(self, test_path):
        mock_metadata_data = """{ 
        no_data
        }""" # Not parsable data
        m = mock_open(read_data=mock_metadata_data)
        
        with patch.object(Path, 'is_dir', side_effect=[True, False]) as mock_is_dir:
            with patch.object(Path, 'iterdir', return_value=[Path]):
                with patch.object(Path, 'exists', return_value=True):
                    with patch('builtins.open', m):
                        with pytest.raises(Exception):
                            local_gen.process_provider(test_path, test_path, test_path, test_path, 0, [], {})
                            mock_is_dir.assert_called()

    """def test_build_markdown(self, test_path, test_string):
        with patch('local_gen.build_markdown')  as mock_markdown:
            local_gen.build_markdown(test_path, test_path, test_path, test_path, test_path)
            mock_markdown.assert_called_once()"""
    
    def test_success(self, test_path, test_string):
        mock_metadata_data = """{ 
        "title": "some title",
        "description": "some description" 
        }""" # Parsable data
        m = mock_open(read_data=mock_metadata_data)
        
        with patch.object(Path, 'is_dir', return_value=True) as mock_is_dir:
            with patch.object(Path, 'iterdir', return_value=[test_path]):
                with patch.object(Path, 'exists', return_value=True):
                    with patch.object(Path, 'mkdir', return_value=True):
                        with patch('builtins.open', m):
                            with patch('local_gen.build_markdown', return_value=MagicMock(name='MockedResult')) as mock_markdown:
                                result = local_gen.process_provider(test_string, test_string, test_path, test_path, 0, [], {})
                                assert result == 1
                                mock_markdown.assert_called_once()
                                mock_is_dir.assert_called()